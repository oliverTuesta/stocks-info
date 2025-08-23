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

    <!-- Table -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                v-for="header in table.getFlatHeaders()"
                :key="header.id"
                :colspan="header.colSpan"
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer hover:bg-gray-100"
                @click="header.column.getCanSort() ? header.column.toggleSorting() : null"
              >
                <div class="flex items-center gap-2">
                  <span>{{ header.isPlaceholder ? null : header.column.columnDef.header }}</span>
                  <div v-if="header.column.getCanSort()" class="flex flex-col">
                    <ChevronUpIcon 
                      class="h-3 w-3 -mb-1" 
                      :class="header.column.getIsSorted() === 'asc' ? 'text-primary' : 'text-gray-400'"
                    />
                    <ChevronDownIcon 
                      class="h-3 w-3" 
                      :class="header.column.getIsSorted() === 'desc' ? 'text-primary' : 'text-gray-400'"
                    />
                  </div>
                </div>
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr
              v-for="row in table.getRowModel().rows"
              :key="row.id"
              class="hover:bg-gray-50"
            >
              <td
                v-for="cell in row.getVisibleCells()"
                :key="cell.id"
                class="px-6 py-4 whitespace-nowrap text-sm text-gray-900"
              >
                <div v-if="cell.column.id === 'ticker'" class="font-mono font-semibold text-primary">
                  {{ cell.getValue() }}
                </div>
                <div v-else-if="cell.column.id === 'company_name'" class="max-w-xs">
                   <div class="font-medium text-gray-900 truncate" :title="cell.getValue()">
                     {{ cell.getValue() }}
                   </div>
                 </div>
                <div v-else-if="cell.column.id === 'industry_name'" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                  {{ cell.getValue() }}
                </div>
                <div v-else-if="cell.column.id === 'exchange'" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                  {{ cell.getValue() }}
                </div>
                <div v-else-if="cell.column.id === 'market_cap'" class="text-right font-medium text-gray-900">
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
        <span>Page {{ table.getState().pagination.pageIndex + 1 }} of {{ table.getPageCount() }}</span>
        <span class="text-gray-500">|</span>
        <span>{{ table.getFilteredRowModel().rows.length }} of {{ totalCompanies }} companies</span>
      </div>

      <div class="flex items-center gap-2">
        <button
          class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="!table.getCanPreviousPage()"
          @click="table.setPageIndex(0)"
        >
          <ChevronLeftIcon class="h-4 w-4" />
        </button>
        <button
          class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="!table.getCanPreviousPage()"
          @click="table.previousPage()"
        >
          <ChevronLeftIcon class="h-4 w-4" />
        </button>
        <button
          class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="!table.getCanNextPage()"
          @click="table.nextPage()"
        >
          <ChevronRightIcon class="h-4 w-4" />
        </button>
        <button
          class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="!table.getCanNextPage()"
          @click="table.setPageIndex(table.getPageCount() - 1)"
        >
          <ChevronRightIcon class="h-4 w-4" />
        </button>
      </div>

      <div class="flex items-center gap-2">
        <span class="text-sm text-gray-700">Show:</span>
        <select
          :value="table.getState().pagination.pageSize"
          @change="(event: Event) => table.setPageSize(Number((event.target as HTMLSelectElement).value))"
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
import { ref, reactive, onMounted, watch } from 'vue'
import {
  type ColumnDef,
  getCoreRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  getFilteredRowModel,
  useVueTable,
  type SortingState,
  type PaginationState,
} from '@tanstack/vue-table'
import {
  SearchIcon,
  ChevronUpIcon,
  ChevronDownIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
} from 'lucide-vue-next'
import { CompanyService } from '@/services/companyService'
import type { Company } from '@/types/company'

// State
const companies = ref<Company[]>([])
const totalCompanies = ref(0)
const loading = ref(false)
const searchQuery = ref('')
const searchTimeout = ref<number | null>(null)

// Table state
const sorting = reactive<SortingState>([])
const pagination = reactive<PaginationState>({
  pageIndex: 0,
  pageSize: 20,
})

// Column definitions
const columns: ColumnDef<Company>[] = [
  {
    accessorKey: 'ticker',
    header: 'Ticker',
    cell: ({ row }) => row.original.ticker,
    enableSorting: true,
  },
  {
    accessorKey: 'company_name',
    header: 'Company Name',
    cell: ({ row }) => row.original.company_name.trim() || row.original.short_name || 'N/A',
    enableSorting: true,
  },
  {
    accessorKey: 'industry_name',
    header: 'Industry',
    cell: ({ row }) => row.original.industry_name || 'N/A',
    enableSorting: true,
  },
  {
    accessorKey: 'exchange',
    header: 'Exchange',
    cell: ({ row }) => row.original.exchange,
    enableSorting: false,
  },
  {
    accessorKey: 'market_cap',
    header: 'Market Cap',
    cell: ({ row }) => `$${(row.original.market_cap / 1000000).toFixed(0)}M`,
    enableSorting: true,
  },
  {
    accessorKey: 'ceo',
    header: 'CEO',
    cell: ({ row }) => row.original.ceo || 'N/A',
    enableSorting: true,
  },
  {
    accessorKey: 'website',
    header: 'Website',
    cell: ({ row }) => row.original.website || null,
    enableSorting: false,
  },
]

// Create table instance
const table = useVueTable({
  get data() {
    return companies.value
  },
  columns,
  state: {
    get sorting() {
      return sorting
    },
    get pagination() {
      return pagination
    },
  },
  onSortingChange: (updater) => {
    if (typeof updater === 'function') {
      Object.assign(sorting, updater(sorting))
    } else {
      Object.assign(sorting, updater)
    }
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
  getSortedRowModel: getSortedRowModel(),
  getFilteredRowModel: getFilteredRowModel(),
  manualSorting: true,
  manualPagination: true,
})

// Fetch companies
const fetchCompanies = async () => {
  loading.value = true
  try {
    const sortBy = sorting.length > 0 ? sorting[0].id : undefined
    const sortOrder = sorting.length > 0 ? sorting[0].desc ? 'desc' : 'asc' : undefined
    
    const response = await CompanyService.getCompanies(
      pagination.pageIndex + 1,
      pagination.pageSize,
      searchQuery.value || undefined,
      sortBy,
      sortOrder
    )
    
    companies.value = response.data
    totalCompanies.value = response.total
  } catch (error) {
    console.error('Failed to fetch companies:', error)
    // In a real app, you'd show a user-friendly error message
  } finally {
    loading.value = false
  }
}

// Handle search with debouncing
const handleSearch = () => {
  if (searchTimeout.value) {
    clearTimeout(searchTimeout.value)
  }
  
  searchTimeout.value = setTimeout(() => {
    pagination.pageIndex = 0
    fetchCompanies()
  }, 300)
}

// Watch for sorting and pagination changes
watch([sorting, pagination], () => {
  fetchCompanies()
}, { deep: true })

// Initial fetch
onMounted(() => {
  fetchCompanies()
})
</script>
