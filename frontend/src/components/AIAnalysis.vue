<template>
  <div class="bg-primary-light rounded-xl border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 bg-primary/10 rounded-full flex items-center justify-center">
          <svg class="w-6 h-6 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
          </svg>
        </div>
        <div>
          <h2 class="text-xl font-semibold text-gray-900">AI Investment Analysis</h2>
          <p class="text-sm text-gray-600">Powered by {{ aiAnalysis?.ai_provider || 'AI' }}</p>
        </div>
      </div>
      
      <!-- Status Badge -->
      <div v-if="aiAnalysis" class="flex items-center gap-2">
        <div class="flex items-center gap-2">
        </div>
        <span class="text-xs text-gray-500">
          {{ formatDate(aiAnalysis.updated_at) }}
        </span>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center py-12">
      <div class="text-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto"></div>
        <p class="mt-4 text-gray-600">Analyzing company data...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="text-center py-8">
      <div class="text-red-500 mb-4">
        <svg class="h-12 w-12 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
        </svg>
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">Analysis Unavailable</h3>
      <p class="text-gray-600 mb-4">{{ error }}</p>
      <button
        @click="fetchAIAnalysis"
        class="inline-flex items-center px-4 py-2 text-sm font-medium rounded-md text-white bg-primary hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary transition-colors"
      >
        Try again
      </button>
    </div>

    <!-- AI Analysis Content -->
    <div v-else-if="aiAnalysis" class="space-y-6">
      <!-- Investment Recommendation -->
      <div class="bg-white rounded-lg p-4 border-l-4 border-primary">
        <h3 class="text-lg font-semibold text-gray-900 mb-3">Investment Recommendation</h3>
        <p class="text-gray-700 leading-relaxed">{{ aiAnalysis.recommendation }}</p>
      </div>

      <!-- Analysis Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Investment Considerations -->
        <div class="bg-white rounded-lg p-4">
          <div class="flex items-center gap-2 mb-3">
            <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900">Investment Considerations</h3>
          </div>
          <p class="text-gray-700 text-sm leading-relaxed">{{ aiAnalysis.investment_considerations }}</p>
        </div>

        <!-- Market Position -->
        <div class="bg-white rounded-lg p-4">
          <div class="flex items-center gap-2 mb-3">
            <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900">Market Position</h3>
          </div>
          <p class="text-gray-700 text-sm leading-relaxed">{{ aiAnalysis.market_position }}</p>
        </div>

        <!-- Recent Events -->
        <div class="bg-white rounded-lg p-4">
          <div class="flex items-center gap-2 mb-3">
            <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900">Recent Events</h3>
          </div>
          <p class="text-gray-700 text-sm leading-relaxed">{{ aiAnalysis.recent_events }}</p>
        </div>

        <!-- Financial Highlights -->
        <div class="bg-white rounded-lg p-4">
          <div class="flex items-center gap-2 mb-3">
            <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900">Financial Highlights</h3>
          </div>
          <p class="text-gray-700 text-sm leading-relaxed">{{ aiAnalysis.financial_highlights }}</p>
        </div>
      </div>

      <!-- Risk Factors & Opportunities -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Risk Factors -->
        <div class="bg-white rounded-lg p-4">
          <div class="flex items-center gap-2 mb-3">
            <div class="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900">Risk Factors</h3>
          </div>
          <p class="text-gray-700 text-sm leading-relaxed">{{ aiAnalysis.risk_factors }}</p>
        </div>

        <!-- Opportunities -->
        <div class="bg-white rounded-lg p-4">
          <div class="flex items-center gap-2 mb-3">
            <div class="w-8 h-8 bg-emerald-100 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900">Growth Opportunities</h3>
          </div>
          <p class="text-gray-700 text-sm leading-relaxed">{{ aiAnalysis.opportunities }}</p>
        </div>
      </div>

      <!-- AI Provider Info -->
      <div class="bg-gray-50 rounded-lg p-4 text-center">
        <p class="text-sm text-gray-600">
          Analysis generated by <span class="font-medium text-gray-900">{{ aiAnalysis.ai_provider }}</span>
        </p>
        <p class="text-xs text-gray-500 mt-1">
          Last updated: {{ formatDate(aiAnalysis.updated_at) }}
        </p>
      </div>
    </div>

    <!-- No Analysis State -->
    <div v-else class="text-center py-8">
      <div class="text-gray-400 mb-4">
        <svg class="h-16 w-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
        </svg>
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">No AI Analysis Available</h3>
      <p class="text-gray-600">AI analysis hasn't been generated for this company yet.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import type { AIAnalysis } from '@/types/aiAnalysis'
import { AIAnalysisService } from '@/services/aiAnalysisService'

// Props
interface Props {
  ticker: string
}

const props = defineProps<Props>()

// Reactive state
const aiAnalysis = ref<AIAnalysis | null>(null)
const loading = ref(false)
const error = ref<string | null>(null)

// Methods
const fetchAIAnalysis = async () => {
  if (!props.ticker) {
    error.value = 'No ticker provided'
    return
  }

  loading.value = true
  error.value = null
  
  try {
    const data = await AIAnalysisService.getCompanyAIAnalysis(props.ticker)
    aiAnalysis.value = data
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Failed to fetch AI analysis'
    console.error('Error fetching AI analysis:', err)
  } finally {
    loading.value = false
  }
}

const formatDate = (dateString: string): string => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

onMounted(() => {
  fetchAIAnalysis()
})

watch(() => props.ticker, () => {
  if (props.ticker) {
    fetchAIAnalysis()
  }
})
</script>


