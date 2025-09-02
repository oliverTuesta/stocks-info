import type { AIAnalysis } from '@/types/aiAnalysis'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

export class AIAnalysisService {
  static async getCompanyAIAnalysis(ticker: string): Promise<AIAnalysis> {
    const response = await fetch(`${API_BASE_URL}/api/companies/ticker/${ticker}/ai-analysis`)
    
    if (!response.ok) {
      throw new Error(`Failed to fetch AI analysis: ${response.statusText}`)
    }

    return await response.json()
  }
}


