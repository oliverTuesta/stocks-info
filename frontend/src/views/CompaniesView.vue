<template>
  <div class="container mx-auto px-4 py-8">
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-2">Companies</h1>
      <p class="text-gray-600">Browse and analyze all available companies</p>
    </div>

    <!-- Search Bar -->
    <div class="mb-6">
      <div class="relative max-w-md">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search companies by name..."
          class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
          @input="handleSearch"
        />
        <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
          <SearchIcon class="h-5 w-5 text-gray-400" />
        </div>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                v-for="header in table.getFlatHeaders()"
                :key="header.id"
                :colspan="header.colSpan"
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                <div class="flex items-center gap-2">
                  <span>{{ header.isPlaceholder ? null : header.column.columnDef.header }}</span>
                </div>
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="row in table.getRowModel().rows" :key="row.id" class="hover:bg-gray-50">
              <td
                v-for="cell in row.getVisibleCells()"
                :key="cell.id"
                class="px-6 py-4 whitespace-nowrap text-sm text-gray-900"
              >
                <div v-if="cell.column.id === 'logo_url'" class="flex items-center justify-center">
                  <img
                    :src="cell.getValue()"
                    :alt="`${row.original.ticker} logo`"
                    class="w-10 h-10 object-contain rounded border-none"
                  />
                </div>
                <div
                  v-else-if="cell.column.id === 'ticker'"
                  class="font-mono font-semibold text-primary"
                >
                  {{ cell.getValue() }}
                </div>
                <div v-else-if="cell.column.id === 'company_name'" class="max-w-40">
                  <div class="font-medium text-gray-900 truncate" :title="cell.getValue()">
                    {{ cell.getValue() }}
                  </div>
                </div>
                <div
                  v-else-if="cell.column.id === 'industry_name'"
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
                >
                  {{ cell.getValue() }}
                </div>
                <div
                  v-else-if="cell.column.id === 'exchange'"
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800"
                >
                  {{ cell.getValue() }}
                </div>
                <div v-else-if="cell.column.id === 'website'">
                  <a
                    v-if="cell.getValue()"
                    :href="cell.getValue()"
                    target="_blank"
                    rel="noopener noreferrer"
                    class="text-primary hover:text-primary-dark underline"
                  >
                    Visit
                  </a>
                  <span v-else class="text-gray-400">N/A</span>
                </div>
                <div v-else class="text-sm text-gray-900">
                  {{ cell.getValue() }}
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex items-center justify-center py-8">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
        <span class="ml-2 text-gray-600">Loading companies...</span>
      </div>

      <!-- Empty State -->
      <div v-else-if="table.getRowModel().rows.length === 0" class="text-center py-8">
        <p class="text-gray-500">No companies found</p>
      </div>
    </div>

    <!-- Pagination -->
    <div class="mt-6 flex items-center justify-between">
      <div class="flex items-center gap-2 text-sm text-gray-700">
        <span>Page {{ currentPage }} of {{ totalPages }}</span>
        <span class="text-gray-500">|</span>
        <span>{{ companies.length }} of {{ totalCompanies }} companies</span>
      </div>

      <div class="flex items-center gap-2">
        <button
          class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="currentPage === 1"
          @click="goToFirstPage"
        >
          <ChevronLeftIcon class="h-4 w-4" />
        </button>
        <button
          class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="currentPage === 1"
          @click="goToPreviousPage"
        >
          <ChevronLeftIcon class="h-4 w-4" />
        </button>
        <button
          class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="currentPage >= totalPages"
          @click="goToNextPage"
        >
          <ChevronRightIcon class="h-4 w-4" />
        </button>
        <button
          class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="currentPage >= totalPages"
          @click="goToLastPage"
        >
          <ChevronRightIcon class="h-4 w-4" />
        </button>
      </div>

      <div class="flex items-center gap-2">
        <span class="text-sm text-gray-700">Show:</span>
        <select
          :value="pagination.pageSize"
          @change="
            (event: Event) => {
              pagination.pageSize = Number((event.target as HTMLSelectElement).value)
              currentPage = 1
              pagination.pageIndex = 0
              fetchCompanies()
            }
          "
          class="px-3 py-2 text-sm border border-gray-300 rounded-md focus:ring-2 focus:ring-primary focus:border-transparent"
        >
          <option v-for="size in [10, 20, 50, 100]" :key="size" :value="size">
            {{ size }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from 'vue'
import {
  type ColumnDef,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  type PaginationState,
  useVueTable,
} from '@tanstack/vue-table'
import { ChevronLeftIcon, ChevronRightIcon, SearchIcon } from 'lucide-vue-next'
import { CompanyService } from '@/services/companyService'
import type { Company } from '@/types/company'

// states
const companies = ref<Company[]>([])
const totalCompanies = ref(0)
const totalPages = ref(0)
const currentPage = ref(1)
const loading = ref(false)
const searchQuery = ref('')
const searchTimeout = ref<number | null>(null)

const pagination = reactive<PaginationState>({
  pageIndex: 0,
  pageSize: 10,
})

const columns: ColumnDef<Company>[] = [
  {
    accessorKey: 'logo_url',
    header: 'Logo',
    cell: ({ row }) => row.original.logo_url,
  },
  {
    accessorKey: 'ticker',
    header: 'Ticker',
    cell: ({ row }) => row.original.ticker,
  },
  {
    accessorKey: 'company_name',
    header: 'Company Name',
    cell: ({ row }) => row.original.company_name.trim() || row.original.short_name || 'N/A',
  },
  {
    accessorKey: 'industry_name',
    header: 'Industry',
    cell: ({ row }) => row.original.industry_name || 'N/A',
  },
  {
    accessorKey: 'exchange',
    header: 'Exchange',
    cell: ({ row }) => row.original.exchange,
  },
  {
    accessorKey: 'website',
    header: 'Website',
    cell: ({ row }) => row.original.website || null,
  },
]

const table = useVueTable({
  get data() {
    return companies.value
  },
  columns,
  state: {
    get pagination() {
      return pagination
    },
  },
  onPaginationChange: (updater) => {
    if (typeof updater === 'function') {
      Object.assign(pagination, updater(pagination))
    } else {
      Object.assign(pagination, updater)
    }
  },
  getCoreRowModel: getCoreRowModel(),
  getPaginationRowModel: getPaginationRowModel(),
  getFilteredRowModel: getFilteredRowModel(),
  manualPagination: true,
  pageCount: computed(() => totalPages.value),
})

const fetchCompanies = async () => {
  loading.value = true
  try {
    const response = await CompanyService.getCompanies(
      currentPage.value,
      pagination.pageSize,
      searchQuery.value || undefined,
    )

    companies.value = response.data
    totalCompanies.value = response.pagination.total
    totalPages.value = response.pagination.total_pages
    currentPage.value = response.pagination.page

    pagination.pageIndex = currentPage.value - 1
  } catch (error) {
    console.error('Failed to fetch companies:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  if (searchTimeout.value) {
    clearTimeout(searchTimeout.value)
  }

  searchTimeout.value = setTimeout(() => {
    currentPage.value = 1
    pagination.pageIndex = 0
    fetchCompanies()
  }, 300)
}

const goToPage = (page: number) => {
  currentPage.value = page
  pagination.pageIndex = page - 1
  fetchCompanies()
}

const goToNextPage = () => {
  if (currentPage.value < totalPages.value) {
    goToPage(currentPage.value + 1)
  }
}

const goToPreviousPage = () => {
  if (currentPage.value > 1) {
    goToPage(currentPage.value - 1)
  }
}

const goToFirstPage = () => {
  goToPage(1)
}

const goToLastPage = () => {
  goToPage(totalPages.value)
}

watch(
  () => pagination.pageSize,
  () => {
    currentPage.value = 1
    pagination.pageIndex = 0
    fetchCompanies()
  },
)

onMounted(() => {
  fetchCompanies()
})
</script>
