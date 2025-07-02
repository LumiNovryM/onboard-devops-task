import { AppSidebar } from "@/components/app-sidebar"
import { DataTable } from "@/components/data-table"
import { SectionCards } from "@/components/section-cards"
import { SiteHeader } from "@/components/site-header"
import {
  SidebarInset,
  SidebarProvider,
} from "@/components/ui/sidebar"
import { useEffect, useState } from 'react'
import axios from 'axios'


interface Message {
  id: number
  sender_id: number
  receiver_id: number
  content: string
  created_at: string
}

interface ApiResponse {
  data: Message[]
}


export default function Page() {
    const [data, setData] = useState<ApiResponse | null>(null)
  
     useEffect(() => {
      axios.get('http://localhost:8080/api/messages')
        .then((res) => setData(res.data))
        .catch((err) => console.error(err))
    }, [])
  return (
    <SidebarProvider
      style={
        {
          "--sidebar-width": "calc(var(--spacing) * 72)",
          "--header-height": "calc(var(--spacing) * 12)",
        } as React.CSSProperties
      }
    >
      <AppSidebar variant="inset" />
      <SidebarInset>
        <SiteHeader />
        <div className="flex flex-1 flex-col">
          <div className="@container/main flex flex-1 flex-col gap-2">
            <div className="flex flex-col gap-4 py-4 md:gap-6 md:py-6">
              <SectionCards />
              <DataTable data={data?.data ?? []} />
            </div>
          </div>
        </div>
      </SidebarInset>
    </SidebarProvider>
  )
}
