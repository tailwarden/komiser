import Head from 'next/head';
import { useRouter } from 'next/router';
import Button from '../components/button/Button';
import EmptyState from '../components/empty-state/EmptyState';
import ErrorPage from '../components/error/ErrorPage';
import InventorySidePanel from '../components/inventory/components/InventorySidePanel';
import InventoryStatsCards from '../components/inventory/components/InventoryStatsCards';
import InventoryTable from '../components/inventory/components/InventoryTable';
import useInventory from '../components/inventory/hooks/useInventory';
import SkeletonInventory from '../components/skeleton/SkeletonInventory';
import SkeletonStats from '../components/skeleton/SkeletonStats';
import Toast from '../components/toast/Toast';

export default function Inventory() {
  const router = useRouter();
  const {
    inventoryStats,
    inventory,
    searchedInventory,
    error,
    query,
    setQuery,
    openModal,
    isOpen,
    closeModal,
    data,
    page,
    goTo,
    tags,
    handleChange,
    addNewTag,
    removeTag,
    loading,
    updateTags,
    toast,
    dismissToast,
    deleteLoading,
    reloadDiv,
    bulkItems,
    onCheckboxChange,
    handleBulkSelection,
    bulkSelectCheckbox,
    openBulkModal,
    updateBulkTags
  } = useInventory();

  return (
    <div className="relative">
      <Head>
        <title>Inventory - Komiser</title>
        <meta name="description" content="Inventory - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <div className="flex gap-6 items-center">
        <p className="text-xl font-medium text-black-900">Inventory</p>
        <div className="relative">
          <Button style="ghost" size="sm">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="none"
              viewBox="0 0 24 24"
            >
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeMiterlimit="10"
                strokeWidth="2"
                d="M5.4 2.1h13.2c1.1 0 2 .9 2 2v2.2c0 .8-.5 1.8-1 2.3l-4.3 3.8c-.6.5-1 1.5-1 2.3V19c0 .6-.4 1.4-.9 1.7l-1.4.9c-1.3.8-3.1-.1-3.1-1.7v-5.3c0-.7-.4-1.6-.8-2.1l-3.8-4c-.5-.5-.9-1.4-.9-2V4.2c0-1.2.9-2.1 2-2.1zM10.93 2.1L6 10"
              ></path>
            </svg>
            Filter by
          </Button>
          <div className="absolute flex flex-col min-w-[16rem] left-0 top-12 bg-white p-4 shadow-xl text-sm rounded-lg z-[11]">
            <Button size="xs" style="ghost" align="left" gap="md">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="none"
                viewBox="0 0 24 24"
              >
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeMiterlimit="10"
                  strokeWidth="2"
                  d="M9 22H7c-4 0-5-1-5-5V7c0-4 1-5 5-5h1.5c1.5 0 1.83.44 2.4 1.2l1.5 2c.38.5.6.8 1.6.8h3c4 0 5 1 5 5v2"
                ></path>
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeMiterlimit="10"
                  strokeWidth="2"
                  d="M13.76 18.32c-2.35.17-2.35 3.57 0 3.74h5.56c.67 0 1.33-.25 1.82-.7 1.65-1.44.77-4.32-1.4-4.59-.78-4.69-7.56-2.91-5.96 1.56"
                ></path>
              </svg>
              Cloud provider
            </Button>
            <Button size="xs" style="ghost" align="left" gap="md">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="none"
                viewBox="0 0 24 24"
              >
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeMiterlimit="10"
                  strokeWidth="2"
                  d="M5.54 11.12c-4.68.33-4.68 7.14 0 7.47h1.92M5.59 11.12C2.38 2.19 15.92-1.38 17.47 8c4.33.55 6.08 6.32 2.8 9.19-1 .91-2.29 1.41-3.64 1.4h-.09"
                ></path>
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeMiterlimit="10"
                  strokeWidth="2"
                  d="M17 16.53c0 .74-.16 1.44-.46 2.06-.08.18-.17.35-.27.51A4.961 4.961 0 0112 21.53c-1.82 0-3.41-.98-4.27-2.43-.1-.16-.19-.33-.27-.51-.3-.62-.46-1.32-.46-2.06 0-2.76 2.24-5 5-5s5 2.24 5 5z"
                ></path>
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M10.44 16.53l.99.99 2.13-1.97"
                ></path>
              </svg>
              Cloud account
            </Button>
            <Button size="xs" style="ghost" align="left" gap="md">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="none"
                viewBox="0 0 24 24"
              >
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M22 12c0-5.52-4.48-10-10-10S2 6.48 2 12s4.48 10 10 10"
                ></path>
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M8 3h1a28.424 28.424 0 000 18H8M15 3c.97 2.92 1.46 5.96 1.46 9"
                ></path>
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M3 16v-1c2.92.97 5.96 1.46 9 1.46M3 9a28.424 28.424 0 0118 0M18.2 21.4a3.2 3.2 0 100-6.4 3.2 3.2 0 000 6.4zM22 22l-1-1"
                ></path>
              </svg>
              Cloud region
            </Button>
            <Button size="xs" style="ghost" align="left" gap="md">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="none"
                viewBox="0 0 24 24"
              >
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeMiterlimit="10"
                  strokeWidth="2"
                  d="M6.37 9.51c-4.08.29-4.07 6.2 0 6.49h9.66c1.17.01 2.3-.43 3.17-1.22 2.86-2.5 1.33-7.5-2.44-7.98C15.41-1.34 3.62 1.75 6.41 9.51M12 16v3M12 23a2 2 0 100-4 2 2 0 000 4zM18 21h-4M10 21H6"
                ></path>
              </svg>
              Cloud service
            </Button>
            <Button size="xs" style="ghost" align="left" gap="md">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="none"
                viewBox="0 0 24 24"
              >
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M4.17 15.3l4.53 4.53a4.78 4.78 0 006.75 0l4.39-4.39a4.78 4.78 0 000-6.75L15.3 4.17a4.75 4.75 0 00-3.6-1.39l-5 .24c-2 .09-3.59 1.68-3.69 3.67l-.24 5c-.06 1.35.45 2.66 1.4 3.61z"
                ></path>
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeWidth="2"
                  d="M9.5 12a2.5 2.5 0 100-5 2.5 2.5 0 000 5z"
                ></path>
              </svg>
              Tags
            </Button>
          </div>
        </div>
      </div>
      <div className="mt-8"></div>

      {/* Toast */}
      {toast && <Toast {...toast} dismissToast={dismissToast} />}

      {/* Error */}
      {((error && !inventoryStats) || (error && !inventory)) && (
        <ErrorPage
          title="Network request error"
          message="There was an error fetching the inventory resources. Check out the server logs for more info and try again."
          action={
            <Button size="lg" style="outline" onClick={() => router.reload()}>
              Refresh the page
            </Button>
          }
        />
      )}

      {/* Empty state */}
      {((!error &&
        inventoryStats &&
        Object.keys(inventoryStats).length === 0) ||
        (!error && inventory && inventory.length === 0)) && (
        <EmptyState
          title="No inventory available"
          message="Check if your connected cloud accounts have active services running or if you have attached the proper permissions."
          action={() => {
            router.push(
              'https://docs.komiser.io/docs/overview/introduction/getting-started/'
            );
          }}
          actionLabel="Check our docs"
          mascotPose="greetings"
        />
      )}

      {/* Inventory stats loading */}
      {!inventoryStats && !error && <SkeletonStats />}

      {/* Inventory stats */}
      <InventoryStatsCards inventoryStats={inventoryStats} error={error} />

      <div className="mt-8"></div>

      {/* Inventory list loading */}
      {!query && !inventory && !error && <SkeletonInventory />}

      {/* Inventory list */}
      <InventoryTable
        error={error}
        inventory={inventory!}
        searchedInventory={searchedInventory!}
        query={query}
        openModal={openModal}
        setQuery={setQuery}
        bulkSelectCheckbox={bulkSelectCheckbox}
        handleBulkSelection={handleBulkSelection}
        bulkItems={bulkItems}
        onCheckboxChange={onCheckboxChange}
        inventoryStats={inventoryStats}
        openBulkModal={openBulkModal}
      />

      {/* Infite scroll trigger */}
      <div ref={reloadDiv}></div>

      {/* Modal */}
      <InventorySidePanel
        isOpen={isOpen}
        closeModal={closeModal}
        data={data!}
        goTo={goTo}
        page={page}
        updateTags={updateTags}
        tags={tags}
        handleChange={handleChange}
        removeTag={removeTag}
        addNewTag={addNewTag}
        loading={loading}
        deleteLoading={deleteLoading}
        bulkItems={bulkItems}
        updateBulkTags={updateBulkTags}
      />
    </div>
  );
}
