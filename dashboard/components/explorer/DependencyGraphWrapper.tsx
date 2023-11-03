import { useRouter } from 'next/router';
import cn from 'classnames';

import { useEffect, useState } from 'react';
import parseURLParams from '@components/inventory/hooks/useInventory/helpers/parseURLParams';
import { InventoryFilterData } from '@components/inventory/hooks/useInventory/types/useInventoryTypes';
import ArrowDownIcon from '@components/icons/ArrowDownIcon';
import EmptyState from '@components/empty-state/EmptyState';
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
  const [filterOpen, setFilterOpen] = useState(false);

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
          <p className="text-lg font-medium text-gray-950">
            Resources Explorer
          </p>
          <div
            className={cn(
              'absolute -top-1 right-24 z-20 flex translate-y-0 cursor-pointer items-center justify-start gap-4 rounded-b-[4px] border-x border-b border-gray-200 bg-white px-4 py-2 text-sm transition',
              { 'translate-y-[105px]': filterOpen }
            )}
            onClick={() => setFilterOpen(!filterOpen)}
          >
            {displayedFilters && displayedFilters?.length > 0 && (
              <span className="bg-darkcyan-100 px-[6px] pb-[3px] pt-[2px] text-xs text-darkcyan-500">
                {displayedFilters?.length}
              </span>
            )}
            <span className="">Filters</span>
            <ArrowDownIcon
              height="16"
              width="16"
              className={cn('transition', {
                'rotate-180': filterOpen
              })}
            />
          </div>
        </div>
        <div
          className={cn(
            'absolute left-0 top-0 z-10 m-0 h-[102px] w-full origin-top scale-y-0 border-b border-gray-200 bg-white px-24 transition',
            { 'scale-y-100': filterOpen }
          )}
        >
          <DependendencyGraphFilter
            router={router}
            hasFilters={hasFilters}
            displayedFilters={displayedFilters}
            deleteFilter={deleteFilter}
          />
        </div>
        {!data?.nodes.length && !data?.edges.length ? (
          <div className="mt-24">
            <EmptyState
              title="We could not find any resources"
              message="It seems like you have no AWS cloud resources associated with your cloud accounts"
              mascotPose="devops"
              secondaryActionLabel="Report an issue"
              actionLabel="Check cloud account"
              secondaryAction={() => {
                router.push(
                  'https://github.com/tailwarden/komiser/issues/new/choose'
                );
              }}
              action={() => {
                router.push('/cloud-accounts');
              }}
            />
          </div>
        ) : (
          <DependencyGraphLoader
            loading={loading}
            data={data}
            error={error}
            fetch={fetch}
          />
        )}
      </div>
    </>
  );
}

export default DependencyGraphWrapper;
