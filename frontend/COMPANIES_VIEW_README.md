# Companies View

This document describes the new Companies view that has been added to the StocksInfo frontend application.

## Features

The Companies view provides a comprehensive table interface for browsing and analyzing all available companies with the following capabilities:

### Table Features
- **Sortable Columns**: Click on column headers to sort by:
  - Ticker (ascending/descending)
  - Company Name (ascending/descending)
  - Market Cap (ascending/descending)
- **Search Functionality**: Search companies by name with real-time filtering
- **Pagination**: Navigate through large datasets with configurable page sizes (10, 20, 50, 100)
- **Responsive Design**: Table adapts to different screen sizes with horizontal scrolling

### Data Display
The table shows the following company information:
- **Ticker**: Company stock symbol
- **Company Name**: Full company name and short name
- **Industry**: Company industry with color-coded badges
- **Exchange**: Stock exchange with color-coded badges
- **Market Cap**: Market capitalization in millions
- **CEO**: Company CEO information
- **Website**: Direct link to company website (if available)

## Technical Implementation

### Technologies Used
- **Vue 3** with Composition API
- **TanStack Table v8** for table functionality
- **Tailwind CSS v4** for styling
- **TypeScript** for type safety

### Key Components
1. **CompaniesView.vue**: Main view component with table implementation
2. **CompanyService**: API service for fetching company data
3. **Company Types**: TypeScript interfaces for company data structure

### API Integration
The view integrates with the backend API endpoint:
```
GET /api/companies
```

Query parameters supported:
- `page`: Page number for pagination
- `page_size`: Number of items per page
- `search`: Search query for company names
- `sort_by`: Column to sort by
- `sort_order`: Sort direction (asc/desc)

## Usage

### Navigation
1. Click on the "Companies" link in the header navigation
2. The view will load with the companies table

### Searching
1. Use the search bar at the top of the page
2. Type company names to filter results
3. Search is debounced (300ms delay) for performance

### Sorting
1. Click on sortable column headers (Ticker, Company Name, Market Cap)
2. Click again to reverse sort order
3. Sort indicators show current sort state

### Pagination
1. Use navigation buttons (<<, <, >, >>) to move between pages
2. Change page size using the dropdown selector
3. Current page and total pages are displayed

## File Structure

```
src/
├── views/
│   └── CompaniesView.vue          # Main companies view
├── services/
│   └── companyService.ts          # API service for companies
├── types/
│   └── company.ts                 # TypeScript interfaces
└── router/
    └── index.ts                   # Route configuration
```

## Configuration

### Environment Variables
Set the following environment variable for the API base URL:
```bash
VITE_API_BASE_URL=http://localhost:8080
```

### Backend Requirements
The backend should implement the `/api/companies` endpoint with:
- Pagination support
- Search functionality
- Sorting capabilities
- Response format matching the `CompaniesResponse` interface

## Styling

The view uses Tailwind CSS v4 with custom color variables:
- Primary colors defined in `src/assets/main.css`
- Responsive design with mobile-first approach
- Hover effects and transitions for better UX

## Performance Features

- **Debounced Search**: Prevents excessive API calls during typing
- **Manual Pagination**: Server-side pagination for large datasets
- **Manual Sorting**: Server-side sorting for better performance
- **Lazy Loading**: Route is lazy-loaded for better initial page load

## Future Enhancements

Potential improvements for future versions:
- Export functionality (CSV, Excel)
- Advanced filtering options
- Company detail modal/page
- Real-time data updates
- Column customization
- Saved search preferences
