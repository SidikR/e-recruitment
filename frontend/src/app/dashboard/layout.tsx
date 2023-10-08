export default function DashboardLayout({
    children, // will be a page or nested layout
  }: {
    children: React.ReactNode
  }) {
    return (
      <section>
        <nav><h2>Ini adalah navbar</h2></nav>
        {children}
      </section>
    )
  }