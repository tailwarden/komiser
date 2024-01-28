import Head from 'next/head';
import Button from '../components/button/Button';
import EmptyState from '../components/empty-state/EmptyState';
import ErrorState from '../components/error-state/ErrorState';
import InventoryActiveFilters from '../components/inventory/components/InventoryActiveFilters';
import InventoryHeader from '../components/inventory/components/InventoryHeader';
import InventoryLayout from '../components/inventory/components/InventoryLayout';
import InventorySidePanel from '../components/inventory/components/InventorySidePanel';
import InventoryStatsCards from '../components/inventory/components/InventoryStatsCards';
import InventoryTable from '../components/inventory/components/InventoryTable';
import InventoryView from '../components/inventory/components/view/InventoryView';
import useInventory from '../components/inventory/hooks/useInventory/useInventory';
import SkeletonFilters from '../components/skeleton/SkeletonFilters';
import SkeletonInventory from '../components/skeleton/SkeletonInventory';
import SkeletonStats from '../components/skeleton/SkeletonStats';
import VerticalSpacing from '../components/spacing/VerticalSpacing';
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
    showToast,
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
    getViews,
    hiddenResources,
    hideResourceFromCustomView,
    hideResourcesLoading,
    setHideOrUnhideHasUpdate,
    hasErrorAndNoInventory,
    hasNoInventory,
    isNotCustomView,
    hasFilterOrCustomView,
    displayFilterIfIsNotCustomView,
    loadingFilters,
    hasFilters,
    loadingInventory,
    isSomeServiceUnavailable
  } = useInventory();

  return (
    <div className="relative">
      <Head>
        <title>Inventory - Komiser</title>
        <meta name="description" content="Inventory - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      {/* Wraps the inventory page and handles the custom views sidebar */}
      <InventoryLayout
        views={views}
        router={router}
        error={error}
        inventory={inventory}
        searchedInventory={searchedInventory}
      >
        <InventoryHeader isNotCustomView={isNotCustomView} />

        {/* Active filters skeleton */}
        {loadingFilters && <SkeletonFilters />}

        {/* Filters bar containing active filters and view button */}
        <div className="relative">
          <InventoryActiveFilters
            hasFilters={hasFilters}
            displayedFilters={displayedFilters}
            isNotCustomView={isNotCustomView}
            deleteFilter={deleteFilter}
            router={router}
          >
            {hasFilterOrCustomView && (
              <InventoryView
                filters={filters}
                displayedFilters={displayedFilters}
                showToast={showToast}
                inventoryStats={inventoryStats}
                router={router}
                views={views}
                getViews={getViews}
                hiddenResources={hiddenResources}
                setHideOrUnhideHasUpdate={setHideOrUnhideHasUpdate}
              />
            )}
          </InventoryActiveFilters>
        </div>
        {/* Inventory stats skeleton */}
        {!error && statsLoading && (
          <SkeletonStats NumOfCards={isNotCustomView ? 3 : 4} />
        )}

        {/* Inventory stats */}
        <InventoryStatsCards
          inventoryStats={inventoryStats}
          isSomeServiceUnavailable={isSomeServiceUnavailable}
          error={error}
          statsLoading={statsLoading}
          hiddenResources={hiddenResources}
        />

        <VerticalSpacing />

        {/* Inventory list skeleton */}
        {loadingInventory && <SkeletonInventory />}

        {/* Inventory list */}
        <InventoryTable
          error={error}
          inventory={inventory}
          searchedInventory={searchedInventory}
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
          hideResourceFromCustomView={hideResourceFromCustomView}
          hideResourcesLoading={hideResourcesLoading}
          showToast={showToast}
        />

        {/* Infite scroll trigger */}
        <div ref={reloadDiv}></div>

        {/* Modal */}
        <InventorySidePanel
          isOpen={isOpen}
          closeModal={closeModal}
          data={data}
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
          tabs={['resource details', 'tags']}
        />

        {/* Error state */}
        {hasErrorAndNoInventory && (
          <ErrorState
            title="Network request error"
            message="There was an error fetching the inventory resources. Check out the server logs for more info and try again."
            action={
              <Button
                size="lg"
                style="secondary"
                onClick={() => router.reload()}
              >
                Refresh the page
              </Button>
            }
          />
        )}

        {/* Empty state */}
        {hasNoInventory && (
          <EmptyState
            title="No inventory available"
            message="Check if your connected cloud accounts have active services running or if you have attached the proper permissions."
            action={() => {
              router.push(
                'https://docs.komiser.io/docs/overview/introduction/getting-started/?utm_source=komiser&utm_medium=referral&utm_campaign=static'
              );
            }}
            actionLabel="Check our docs"
            mascotPose="thinking"
          />
        )}
      </InventoryLayout>
    </div>
  );
}
