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
    <>
      <Head>
        <title>Inventory - Komiser</title>
        <meta name="description" content="Inventory - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <p className="text-xl font-medium text-black-900">Inventory</p>
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
    </>
  );
}
