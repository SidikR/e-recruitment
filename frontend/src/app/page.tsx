import Image from 'next/image'
import RootLayout from './layout'
import DashboardLayout from './dashboard/layout'
import Dashboard from './dashboard/page'

export default function Home() {
  return (
    <RootLayout>
      <Dashboard/>
    </RootLayout>
  )
}
