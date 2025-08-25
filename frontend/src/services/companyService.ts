import type { Company, CompaniesResponse } from '@/types/company'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

export class CompanyService {
  static async getCompanies(
    page: number = 1,
    pageSize: number = 20,
    search?: string
  ): Promise<CompaniesResponse> {
    const params = new URLSearchParams({
      page: page.toString(),
      page_size: pageSize.toString(),
    })

    if (search) {
      params.append('search', search)
    }

    const response = await fetch(`${API_BASE_URL}/api/companies?${params}`)
    
    if (!response.ok) {
      throw new Error(`Failed to fetch companies: ${response.statusText}`)
    }

    let data = await response.json()
    console.log(data)
    return data;
  }

  static async getAllCompanies(): Promise<Company[]> {
    const response = await fetch(`${API_BASE_URL}/api/companies`)
    
    if (!response.ok) {
      throw new Error(`Failed to fetch companies: ${response.statusText}`)
    }

    const data = await response.json()
    return data.data || data
  }
}
