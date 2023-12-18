import { ReactNode, useState } from 'react';
import { NextRouter } from 'next/router';
import useFilterWizard from '@components/inventory/components/filter/hooks/useFilterWizard';
import useInventory from '@components/inventory/hooks/useInventory/useInventory';
import { InventoryFilterData } from '@components/inventory/hooks/useInventory/types/useInventoryTypes';
import PlusIcon from '@components/icons/PlusIcon';
import Button from '@components/button/Button';
import CloseIcon from '@components/icons/CloseIcon';
import DependencyGraphFilterSummary from './DependencyGraphFilterSummary';
import DependencyGraphFilterDropdown from './DependencyGraphFilterDropdown';

type DependendencyGraphFilterProps = {
  hasFilters: boolean | undefined;
  displayedFilters: InventoryFilterData[] | undefined;
  deleteFilter: (idx: number) => void;
  router: NextRouter;
  children?: ReactNode;
};

function DependendencyGraphFilter({
  hasFilters,
  displayedFilters,
  deleteFilter,
  router,
  children
}: DependendencyGraphFilterProps) {
  const [skippedSearch, setSkippedSearch] = useState(0);
  const { toggle, isOpen } = useFilterWizard({ router, setSkippedSearch });

  return (
    <div className="flex h-full items-center justify-between rounded-lg bg-white px-6 py-2">
      {!hasFilters ? (
        <>
          <div
            className="flex w-fit cursor-pointer items-center gap-1 overflow-hidden rounded-md border-2 border-dashed border-gray-300 border-opacity-60 px-3 py-1"
            onClick={toggle}
          >
            <PlusIcon width={16} height={16} />
            <span className="font-sans text-sm text-gray-700">Filter</span>
          </div>
          {isOpen && (
            <DependencyGraphFilterDropdown
              position={'top-10'}
              toggle={toggle}
              closeDropdownAfterAdd={true}
            />
          )}
        </>
      ) : (
        <div className="flex flex-wrap items-center gap-x-4 gap-y-2">
          <div className="h-full text-sm text-gray-700">Filters</div>
          {displayedFilters &&
            displayedFilters.map((activeFilter, idx) => (
              <DependencyGraphFilterSummary
                key={idx}
                id={idx}
                data={activeFilter}
                deleteFilter={deleteFilter}
              />
            ))}

          <div className="flex items-center gap-3">
            <div className="relative">
              <div className="cursor-pointer" onClick={toggle}>
                <PlusIcon className="h-6 w-6 rounded-full border-dashed border-gray-300 p-1 hover:border hover:bg-gray-700 hover:bg-opacity-10" />
              </div>
              {isOpen && (
                <DependencyGraphFilterDropdown
                  position={'top-10'}
                  toggle={toggle}
                  closeDropdownAfterAdd={false}
                />
              )}
            </div>

            <div className="border-x-1 h-7 border"></div>
            <div
              className="flex cursor-pointer items-center gap-1 text-gray-700 hover:text-gray-950"
              onClick={() => router.push(router.pathname)}
            >
              <span className="font-sans text-[14px] font-semibold ">
                Clear filters
              </span>
              <CloseIcon className="h-5 w-5 opacity-70" />
            </div>
          </div>
        </div>
      )}

      {children}
    </div>
  );
}

export default DependendencyGraphFilter;
