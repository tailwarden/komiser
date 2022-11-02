import Head from 'next/head'

export default function Home() {
  return (
    <div>
      <table className="table-auto text-sm text-left bg-white dark:bg-purplin-700 text-gray-900 dark:text-black-100 w-full">
        <thead className="sticky top-0 bg-white dark:bg-purplin-700 z-20">
          <tr>
            <th>Name</th>
            <th className="py-4 pl-4 pr-6">Cloud</th>
            <th className="py-4 px-6">Service</th>
            <th className="py-4 px-6">Name</th>
            <th className="py-4 px-6">Region</th>
            <th className="py-4 px-6">Account</th>
            <th className="py-4 px-6">Cost</th>
            <th className="py-4 px-6"></th>
          </tr>
        </thead>
        <tbody>

        </tbody>
      </table>
    </div>
  )
}
