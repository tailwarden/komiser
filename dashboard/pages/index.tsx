import Head from 'next/head';
import Button from '../components/button/Button';
import EmptyState from '../components/empty-state/EmptyState';
import ErrorPage from '../components/error/ErrorPage';
import InventoryFilter from '../components/inventory/components/filter/InventoryFilter';
import InventoryFilterSummary from '../components/inventory/components/filter/InventoryFilterSummary';
import InventorySidePanel from '../components/inventory/components/InventorySidePanel';
import InventoryStatsCards from '../components/inventory/components/InventoryStatsCards';
import InventoryTable from '../components/inventory/components/InventoryTable';
import InventoryViewsTabs from '../components/inventory/components/InventoryViewsTabs';
import InventoryView from '../components/inventory/components/view/InventoryView';
import useInventory from '../components/inventory/hooks/useInventory';
import SkeletonFilters from '../components/skeleton/SkeletonFilters';
import SkeletonInventory from '../components/skeleton/SkeletonInventory';
import SkeletonStats from '../components/skeleton/SkeletonStats';
import Toast from '../components/toast/Toast';

export default function Inventory() {
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
    setToast,
    dismissToast,
    deleteLoading,
    reloadDiv,
    bulkItems,
    onCheckboxChange,
    handleBulkSelection,
    bulkSelectCheckbox,
    openBulkModal,
    updateBulkTags,
    router,
    filters,
    displayedFilters,
    setSkippedSearch,
    deleteFilter,
    searchedLoading,
    statsLoading,
    views,
    getViews
  } = useInventory();

  return (
    <div className="relative">
      <Head>
        <title>Inventory - Komiser</title>
        <meta name="description" content="Inventory - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className="flex items-center justify-between gap-8">
        {/* Inventory views */}
        <InventoryViewsTabs views={views} router={router} />

        <div className="flex items-center gap-4 flex-shrink-0">
          {/* Save/update views button */}
          {((filters && filters.length > 0) || router.query.view) && (
            <InventoryView
              filters={filters!}
              displayedFilters={displayedFilters!}
              setToast={setToast}
              inventoryStats={inventoryStats!}
              router={router}
              views={views}
              getViews={getViews}
            />
          )}

          {/* Filter by dropdown */}
          {!error &&
            ((inventory && inventory.length > 0) ||
              (searchedInventory && searchedInventory.length > 0)) &&
            !router.query.view && (
              <InventoryFilter
                router={router}
                setSkippedSearch={setSkippedSearch}
                setToast={setToast}
              />
            )}
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

      {/* Filters skeleton */}
      {Object.keys(router.query).length > 0 && !displayedFilters && (
        <SkeletonFilters />
      )}

      {/* Active filters list */}
      {Object.keys(router.query).length > 0 &&
        displayedFilters &&
        displayedFilters.length > 0 && (
          <div className="flex flex-wrap items-center gap-x-4 gap-y-2 bg-white py-2 px-6 rounded-lg mb-8">
            <div className="text-sm text-black-400 h-full">Filters</div>
            {displayedFilters.map((activeFilter, idx) => (
              <InventoryFilterSummary
                key={idx}
                id={idx}
                data={activeFilter}
                deleteFilter={!router.query.view ? deleteFilter : undefined}
              />
            ))}
            {!router.query.view && (
              <Button size="sm" style="ghost" onClick={() => router.push('/')}>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="none"
                  viewBox="0 0 24 24"
                  className="text-black-400"
                >
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeMiterlimit="10"
                    strokeWidth="2"
                    d="M13.41 20.79L12 21.7c-1.31.81-3.13-.1-3.13-1.72v-5.35c0-.71-.4-1.62-.81-2.12L4.22 8.47c-.51-.51-.91-1.41-.91-2.02V4.13c0-1.21.91-2.12 2.02-2.12h13.34c1.11 0 2.02.91 2.02 2.02v2.22c0 .81-.51 1.82-1.01 2.32"
                  ></path>
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeMiterlimit="10"
                    strokeWidth="2"
                    d="M21.63 14.75c0 .89-.25 1.73-.69 2.45a4.709 4.709 0 01-4.06 2.3 4.73 4.73 0 01-4.06-2.3 4.66 4.66 0 01-.69-2.45c0-2.62 2.13-4.75 4.75-4.75s4.75 2.13 4.75 4.75zM18.15 15.99l-2.51-2.51M18.13 13.51l-2.51 2.51"
                  ></path>
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeMiterlimit="10"
                    strokeWidth="2"
                    d="M20.69 4.02v2.22c0 .81-.51 1.82-1.01 2.33l-1.76 1.55a4.42 4.42 0 00-1.04-.12c-2.62 0-4.75 2.13-4.75 4.75 0 .89.25 1.73.69 2.45.37.62.88 1.15 1.5 1.53v.34c0 .61-.4 1.42-.91 1.72L12 21.7c-1.31.81-3.13-.1-3.13-1.72v-5.35c0-.71-.41-1.62-.81-2.12L4.22 8.47c-.5-.51-.91-1.42-.91-2.02V4.12C3.31 2.91 4.22 2 5.33 2h13.34c1.11 0 2.02.91 2.02 2.02z"
                  ></path>
                </svg>
                Clear filters
              </Button>
            )}
          </div>
        )}
      {/* Inventory stats loading */}
      {!error && statsLoading && <SkeletonStats />}

      {/* Inventory stats */}
      <InventoryStatsCards
        inventoryStats={inventoryStats}
        error={error}
        statsLoading={statsLoading}
      />

      <div className="mt-8"></div>
      {/* Inventory list loading */}
      {!inventory &&
        !error &&
        !query &&
        !displayedFilters &&
        !router.query.view && <SkeletonInventory />}

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
        router={router}
        searchedLoading={searchedLoading}
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
