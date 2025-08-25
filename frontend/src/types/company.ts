export interface Company {
  id: number
  ticker: string
  company_name: string
  short_name: string
  industry_name: string | null
  description: string | null
  website: string | null
  logo_url: string | null
  ceo: string | null
  exchange: string
  market_cap: number
  created_at: string
  updated_at: string
}

export interface PaginationInfo {
  page: number
  page_size: number
  total: number
  total_pages: number
}

export interface CompaniesResponse {
  data: Company[]
  pagination: PaginationInfo
}
