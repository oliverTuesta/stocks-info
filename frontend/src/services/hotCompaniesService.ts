import type { Company } from '@/types/company'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

export class HotCompaniesService {
  static async getHotCompanies(limit: number = 10): Promise<Company[]> {
    const response = await fetch(`${API_BASE_URL}/api/companies/hot?limit=${limit}`)
    if (!response.ok) {
      throw new Error(`Failed to fetch hot companies: ${response.statusText}`)
    }

    return await response.json()
  }
}
