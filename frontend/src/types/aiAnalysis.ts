export interface AIAnalysis {
  id: number
  company_id: number
  investment_considerations: string
  recent_events: string
  market_position: string
  risk_factors: string
  opportunities: string
  financial_highlights: string
  recommendation: string
  ai_provider: string
  status: 'completed' | 'pending' | 'failed'
  created_at: string
  updated_at: string
}


