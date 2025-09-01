<template>
  <div>
    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center min-h-screen">
      <div class="text-center">
        <div class="animate-spin rounded-full h-16 w-16 border-b-2 border-primary mx-auto"></div>
        <p class="mt-6 text-xl text-gray-600">Loading company details...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="min-h-screen flex items-center justify-center">
      <div class="text-center max-w-md mx-auto px-4">
        <div class="text-red-500 mb-6">
          <svg class="h-20 w-20 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"
            />
          </svg>
        </div>
        <h1 class="text-2xl font-bold text-gray-900 mb-4">Something went wrong</h1>
        <p class="text-gray-600 mb-6">{{ error }}</p>
        <div class="space-y-3">
          <button
            @click="fetchCompanyDetails"
            class="w-full inline-flex items-center justify-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-white bg-primary hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary transition-colors"
          >
            Try again
          </button>
          <router-link
            to="/companies"
            class="w-full inline-flex items-center justify-center px-6 py-3 border border-gray-300 text-base font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary transition-colors"
          >
            Back to Companies
          </router-link>
        </div>
      </div>
    </div>

    <!-- Company Details -->
    <div v-else-if="company" class="container mx-auto px-4 py-8">
      <!-- Back Button -->
      <div class="mb-6">
        <router-link
          to="/companies"
          class="inline-flex items-center gap-2 text-gray-600 hover:text-gray-900 transition-colors"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
          Back to Companies
        </router-link>
      </div>

      <div class="rounded-lg p-6 mb-6 bg-primary-light">
        <div class="flex flex-col lg:flex-row lg:items-center gap-6">
          <div class="flex gap-6 w-full">
            <div class="w-24 h-24 rounded-xl bg-white flex items-center justify-center overflow-hidden flex-shrink-0">
              <img
                v-if="company.logo_url"
                :src="company.logo_url"
                :alt="company.short_name"
                class="w-full h-full object-contain"
                @error="handleImageError"
              />
              <div
                v-else
                class="w-full h-full bg-gradient-to-br from-blue-100 to-blue-200 flex items-center justify-center text-blue-600 font-bold text-4xl"
              >
                {{ company.ticker.charAt(0) }}
              </div>
            </div>

            <div class="w-full">
              <div class="flex justify-between">
                <div class="flex items-center gap-3 mb-2">
                  <h1>{{ company.short_name || company.company_name }}</h1>
                  <span class="px-3 py-1 bg-primary text-white text-sm font-semibold rounded-full">
                  {{ company.ticker }}
                </span>
                </div>
                <div class="flex flex-col sm:flex-row gap-3">
                  <a
                    v-if="company.website"
                    :href="company.website"
                    target="_blank"
                    rel="noopener noreferrer"
                    class="inline-flex items-center justify-center px-6 py-3 text-base font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary transition-colors"
                  >
                    <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                    </svg>
                    Visit Website
                  </a>
                </div>
              </div>
              <p class="text-lg text-gray-600 mb-4">{{ company.company_name }}</p>
              
              <!-- Quick Stats -->
              <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                <div>
                  <p class="text-sm text-gray-500">Sector</p>
                  <p class="font-semibold text-gray-900">{{ company.sector || 'N/A' }}</p>
                </div>
                <div>
                  <p class="text-sm text-gray-500">Industry</p>
                  <p class="font-semibold text-gray-900">{{ company.industry_name || 'N/A' }}</p>
                </div>
                <div>
                  <p class="text-sm text-gray-500">Exchange</p>
                  <p class="font-semibold text-gray-900">{{ company.exchange }}</p>
                </div>
                <div>
                  <p class="text-sm text-gray-500">Market Cap</p>
                  <p class="font-semibold text-gray-900">{{ formatMarketCap(company.market_cap) }}</p>
                </div>

              </div>
            </div>
          </div>

          <!-- Action Buttons -->

        </div>
      </div>

      <!-- Company Description -->
      <div class="rounded-xl bg-primary-light p-6 mb-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">About {{ company.short_name || company.company_name }}</h2>
        <p class="text-gray-700 leading-relaxed">{{ company.description || 'No description available.' }}</p>
        
        <div class="mt-6 pt-6 border-t border-gray-200">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-3">Company Information</h3>
              <div class="space-y-3">
                <div class="flex justify-between">
                  <span class="text-gray-600">CEO:</span>
                  <span class="font-medium text-gray-900">{{ company.ceo || 'N/A' }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-gray-600">Last Updated:</span>
                  <span class="font-medium text-gray-900">{{ formatDate(company.updated_at) }}</span>
                </div>
              </div>
            </div>
            
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-3">Financial Overview</h3>
              <div class="space-y-3">
                <div class="flex justify-between">
                  <span class="text-gray-600">Market Cap:</span>
                  <span class="font-medium text-gray-900">{{ formatMarketCap(company.market_cap) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-gray-600">Exchange:</span>
                  <span class="font-medium text-gray-900">{{ company.exchange }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-gray-600">Sector:</span>
                  <span class="font-medium text-gray-900">{{ company.sector || 'N/A' }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Latest Analysis -->
      <div v-if="company.analyses && company.analyses.length > 0" class="bg-primary-light rounded-xl border border-gray-200 p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">Latest Analyst Analysis</h2>
        
        <div class="space-y-4">
          <div
            v-for="analysis in company.analyses"
            :key="analysis.id"
            class="rounded-lg p-4 bg-white"
          >
            <div class="flex items-start justify-between mb-3">
              <div class="flex items-center gap-3">
                <div class="w-12 h-12 bg-primary-light rounded-full flex items-center justify-center">
                  <ChartLineIcon class="w-8 h-8"/>
                </div>
                <div>
                  <h3 class="font-semibold text-gray-900">{{ analysis.brokerage }}</h3>
                  <p class="text-sm text-gray-600">{{ formatDate(analysis.time) }}</p>
                </div>
              </div>
              
              <!-- Rating Change -->
              <div class="text-right">
                <div class="flex items-center gap-2 mb-1">
                  <span class="text-sm text-gray-500">Rating:</span>
                  <span class="px-2 py-1 text-xs font-medium rounded-full" :class="getRatingClass(analysis.rating_to)">
                    {{ analysis.rating_to }}
                  </span>
                </div>
                <div v-if="analysis.rating_from !== analysis.rating_to" class="text-xs text-gray-500">
                  {{ analysis.rating_from }} → {{ analysis.rating_to }}
                </div>
              </div>
            </div>

            <!-- Analysis Details -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div>
                <p class="text-sm text-gray-500 mb-1">Action</p>
                <p class="font-medium text-gray-900">{{ analysis.action }}</p>
              </div>
              <div>
                <p class="text-sm text-gray-500 mb-1">Previous Target</p>
                <p class="font-medium text-gray-900">${{ analysis.target_from }}</p>
              </div>
              <div>
                <p class="text-sm text-gray-500 mb-1">New Target</p>
                <p class="font-medium text-gray-900" :class="getTargetChangeClass(analysis.target_from, analysis.target_to)">
                  ${{ analysis.target_to }}
                </p>
              </div>
            </div>

            <!-- Target Change Summary -->
            <div class="mt-3 pt-3 border-t border-gray-200">
              <div class="flex items-center gap-2">
                <span class="text-sm text-gray-500">Target Change:</span>
                <span class="font-semibold" :class="getTargetChangeClass(analysis.target_from, analysis.target_to)">
                  {{ getTargetChangeText(analysis.target_from, analysis.target_to) }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- No Analysis State -->
      <div v-else class="rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="text-center py-8">
          <div class="text-gray-400 mb-4">
            <svg class="h-16 w-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
          <h3 class="text-lg font-medium text-gray-900 mb-2">No Analysis Available</h3>
          <p class="text-gray-600">This company doesn't have any analyst coverage yet.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { Company } from '@/types/company'
import { CompanyDetailsService } from '@/services/companyDetailsService'
import { ChartLineIcon } from 'lucide-vue-next'

// Route and router
const route = useRoute()
const router = useRouter()

// Reactive state
const company = ref<Company | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)

// Computed properties
const ticker = computed(() => route.params.ticker as string)

// Methods
const fetchCompanyDetails = async () => {
  if (!ticker.value) {
    error.value = 'No ticker provided'
    loading.value = false
    return
  }

  loading.value = true
  error.value = null
  
  try {
    const data = await CompanyDetailsService.getCompanyByTicker(ticker.value)
    company.value = data
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'An unexpected error occurred'
    console.error('Error fetching company details:', err)
  } finally {
    loading.value = false
  }
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

const formatDate = (dateString: string): string => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const getRatingClass = (rating: string): string => {
  const ratingLower = rating.toLowerCase()
  if (ratingLower.includes('buy') || ratingLower.includes('strong buy')) {
    return 'bg-green-100 text-green-800'
  } else if (ratingLower.includes('sell') || ratingLower.includes('strong sell')) {
    return 'bg-red-100 text-red-800'
  } else if (ratingLower.includes('hold') || ratingLower.includes('neutral')) {
    return 'bg-yellow-100 text-yellow-800'
  }
  return 'bg-gray-100 text-gray-800'
}

const getTargetChangeClass = (from: number, to: number): string => {
  if (to > from) return 'text-green-600'
  if (to < from) return 'text-red-600'
  return 'text-gray-600'
}

const getTargetChangeText = (from: number, to: number): string => {
  const change = to - from
  const percentage = ((change / from) * 100).toFixed(1)
  
  if (change > 0) {
    return `+$${change.toLocaleString()} (+${percentage}%)`
  } else if (change < 0) {
    return `-$${Math.abs(change).toLocaleString()} (-${percentage}%)`
  }
  return 'No change'
}

onMounted(() => {
  fetchCompanyDetails()
})

// Watch for route changes
watch(() => route.params.ticker, () => {
  if (route.params.ticker) {
    fetchCompanyDetails()
  }
})
</script>
