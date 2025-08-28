import type { Company } from '@/types/company'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

export class CompanyDetailsService {
  static async getCompanyByTicker(ticker: string): Promise<Company> {
    const response = await fetch(`${API_BASE_URL}/api/companies/ticker/${ticker}`)
    
    if (!response.ok) {
      throw new Error(`Failed to fetch company details: ${response.statusText}`)
    }

    return await response.json()
  }
}
