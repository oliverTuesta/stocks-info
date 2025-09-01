<template>
  <div>
    <!-- Header -->
    <div>
      <h2 class="my-auto flex items-center">
        <FlameIcon class="w-10 h-10 mr-2 text-rose-500"/>
        Hot Companies
      </h2>
      <p>Top 10 trending companies with recent analyst activity</p>
    </div>

    <div class="p-6 bg-primary-light mt-6 rounded-xl">
      <!-- Loading State -->
      <div v-if="loading && companies.length === 0" class="flex items-center justify-center py-12">
        <div class="text-center">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto"></div>
          <p class="mt-4 text-gray-600">Loading hot companies...</p>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-12 bg-white rounded-lg">
        <div class="text-red-500 mb-4">
          <TriangleAlertIcon class="text-red w-12 h-12 mx-auto"  stroke="currentColor"/>
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-2">Something went wrong</h3>
        <p class="text-gray-600 mb-4">{{ error }}</p>
        <button
          @click="fetchHotCompanies"
          class="inline-flex items-center px-4 py-2 text-sm font-medium rounded-md text-white bg-primary hover:bg-primary-light focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary"
        >
          Try again
        </button>
      </div>

      <!-- Companies Grid -->
      <div v-else-if="companies.length > 0" class="grid grid-cols-1 md:grid-cols-2  gap-4">
        <div
          v-for="(company, index) in companies"
          :key="company.id"
          class="group relative bg-white rounded-lg p-4 hover:shadow transition-all duration-200 cursor-pointer"
          @click="handleCompanyClick(company)"
        >
          <!-- Rank Badge -->
          <div class="absolute -top-2 -left-2 w-8 h-8 rounded-full bg-primary flex items-center justify-center text-white text-sm font-bold shadow-lg">
            {{ index + 1 }}
          </div>

          <!-- Company Logo -->
          <div class="flex items-center gap-3 mb-3">
            <div class="w-12 h-12 rounded-lg bg-white border border-gray-200 flex items-center justify-center overflow-hidden">
              <img
                v-if="company.logo_url"
                :src="company.logo_url"
                :alt="company.short_name"
                class="w-full h-full object-contain"
                @error="handleImageError"
              />
              <div
                v-else
                class="w-full h-full bg-gradient-to-br from-blue-100 to-blue-200 flex items-center justify-center text-blue-600 font-bold text-lg"
              >
                {{ company.ticker.charAt(0) }}
              </div>
            </div>
            <div class="flex-1 min-w-0">
              <h3 class="font-semibold text-gray-900 truncate">{{ company.short_name || company.company_name }}</h3>
              <p class="text-sm text-gray-600 font-mono">{{ company.ticker }}</p>
            </div>
          </div>

          <!-- Company Info -->
          <div class="space-y-2 text-sm">
            <div class="flex items-center gap-2">
              <span class="text-gray-500">Sector:</span>
              <span class="font-medium text-gray-700 truncate">{{ company.sector || 'N/A' }}</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-gray-500">Market Cap:</span>
              <span class="font-medium text-gray-700">{{ formatMarketCap(company.market_cap) }}</span>
            </div>
          </div>

          <!-- Analysis Info -->
          <div v-if="company.analyses && company.analyses.length > 0" class="mt-3 pt-3 border-t border-gray-200">
            <div class="flex items-center gap-2 text-xs">
              <span class="text-gray-500">Latest:</span>
              <span class="font-medium text-gray-700">{{ getLatestAnalysis(company.analyses) }}</span>
            </div>
          </div>

          <!-- Hover Effect -->
          <div class="absolute inset-0 rounded-lg bg-gradient-to-r from-primary-light/5 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-200 pointer-events-none"></div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="text-center py-12">
        <div class="text-gray-400 mb-4">
          <svg class="h-12 w-12 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"
            />
          </svg>
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-2">No hot companies found</h3>
        <p class="text-gray-600">Try refreshing the data or check back later.</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import type { Company } from '@/types/company'
import { HotCompaniesService } from '@/services/hotCompaniesService'
import { FlameIcon, TriangleAlertIcon } from 'lucide-vue-next'

interface Props {
  baseURL?: string
}

const props = withDefaults(defineProps<Props>(), {
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'
})

const emit = defineEmits<{
  companyClick: [company: Company]
}>()

const companies = ref<Company[]>([])
const loading = ref(false)
const error = ref<string | null>(null)

const fetchHotCompanies = async () => {
  loading.value = true
  error.value = null
  
  try {
    const data = await HotCompaniesService.getHotCompanies()
    companies.value = data
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'An unexpected error occurred'
    console.error('Error fetching hot companies:', err)
  } finally {
    loading.value = false
  }
}

const handleCompanyClick = (company: Company) => {
  router.push(`/companies/${company.ticker}`)
}

const handleImageError = (event: Event) => {
  const target = event.target as HTMLImageElement
  target.style.display = 'none'
  target.nextElementSibling?.classList.remove('hidden')
}

const formatMarketCap = (marketCap: number): string => {
  if (marketCap >= 1e12) {
    return `$${(marketCap / 1e12).toFixed(2)}T`
  } else if (marketCap >= 1e9) {
    return `$${(marketCap / 1e9).toFixed(2)}B`
  } else if (marketCap >= 1e6) {
    return `$${(marketCap / 1e6).toFixed(2)}M`
  } else if (marketCap >= 1e3) {
    return `$${(marketCap / 1e3).toFixed(2)}K`
  }
  return `$${marketCap.toFixed(2)}`
}

const getLatestAnalysis = (analyses: any[]): string => {
  if (!analyses || analyses.length === 0) return 'No analysis'
  
  const latest = analyses[0]
  const targetChange = latest.target_to - latest.target_from
  const percentageChange = ((targetChange / latest.target_from) * 100).toFixed(1)
  const isPositive = targetChange > 0
  
  let icon = '😶‍🌫️'
  if (isPositive) icon = '🔥'
  if (targetChange < 0) icon = '💤'
  
  return `${icon} ${isPositive ? '+' : ''}${percentageChange}% | ${latest.rating_to}`
}

onMounted(() => {
  fetchHotCompanies()
})
</script>
