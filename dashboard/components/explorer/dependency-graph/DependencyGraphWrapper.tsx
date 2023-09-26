import { useRouter } from 'next/router';

import { useEffect } from 'react';
import parseURLParams from '@components/inventory/hooks/useInventory/helpers/parseURLParams';
import { InventoryFilterData } from '@components/inventory/hooks/useInventory/types/useInventoryTypes';
import DependencyGraphLoader from './DependencyGraphLoader';
import DependendencyGraphFilter from './filter/DependendencyGraphFilter';
import useDependencyGraph from './hooks/useDependencyGraph';

function DependencyGraphWrapper() {
  const {
    loading,
    data,
    error,
    fetch,
    filters,
    displayedFilters,
    setDisplayedFilters,
    deleteFilter,
    setFilters
  } = useDependencyGraph();
  const router = useRouter();

  useEffect(() => {
    const newFilters: InventoryFilterData[] = Object.keys(router.query).map(
      param => parseURLParams(param as string, 'fetch')
    );
    const newFiltersToDisplay: InventoryFilterData[] = Object.keys(
      router.query
    ).map(param => parseURLParams(param as string, 'display'));

    setFilters(newFilters);
    setDisplayedFilters(newFiltersToDisplay);
  }, [router.query]);

  useEffect(() => {
    const newFilters: InventoryFilterData[] = Object.keys(router.query).map(
      param => parseURLParams(param as string, 'fetch')
    );
    const newFiltersToDisplay: InventoryFilterData[] = Object.keys(
      router.query
    ).map(param => parseURLParams(param as string, 'display'));

    setFilters(newFilters);
    setDisplayedFilters(newFiltersToDisplay);
  }, []);

  const hasFilters =
    Object.keys(router.query).length > 0 &&
    displayedFilters &&
    displayedFilters.length > 0;

  return (
    <>
      <div className="flex h-[calc(100vh-145px)] w-full flex-col">
        <div className="flex flex-row justify-between gap-2">
          <p className="text-lg font-medium text-black-900">Graph View</p>
          <div className="absolute -top-1 right-24 border-x border-b border-black-170 bg-white p-2 text-sm">
            Filters
          </div>
        </div>
        <div>
          <DependendencyGraphFilter
            router={router}
            hasFilters={hasFilters}
            displayedFilters={displayedFilters}
            deleteFilter={deleteFilter}
          />
        </div>
        <DependencyGraphLoader
          loading={loading}
          data={data}
          error={error}
          fetch={fetch}
        />
      </div>
    </>
  );
}

export default DependencyGraphWrapper;
