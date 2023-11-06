import { useEffect } from 'react';
import useInventory from '../hooks/useInventory/useInventory';
import useFilterWizard from './filter/hooks/useFilterWizard';
import Button from '../../button/Button';
import InventoryFilterBreadcrumbs from './filter/InventoryFilterBreadcrumbs';
import InventoryFilterSummary from './filter/InventoryFilterSummary';
import InventoryFilterField from './filter/InventoryFilterField';
import InventoryFilterOperator from './filter/InventoryFilterOperator';
import InventoryFilterValue from './filter/InventoryFilterValue';

type InventoryFilterDropdownProps = {
  position: string;
  closeDropdownAfterAdd: boolean;
  toggle: () => void;
};

export default function InventoryFilterDropdown({
  position,
  toggle,
  closeDropdownAfterAdd
}: InventoryFilterDropdownProps) {
  const { setSkippedSearch, router, showToast } = useInventory();

  const {
    // toggle,
    step,
    goTo,
    handleField,
    handleOperator,
    handleTagKey,
    handleValueCheck,
    handleValueInput,
    costBetween,
    handleCostBetween,
    inlineError,
    data,
    resetData,
    cleanValues,
    filter
  } = useFilterWizard({ router, setSkippedSearch });

  useEffect(() => {
    cleanValues();
  }, []);

  return (
    <>
      {/* Dropdown transparent backdrop */}
      <div
        onClick={toggle}
        className="fixed left-0 top-0 z-20 hidden h-screen w-screen animate-fade-in bg-transparent opacity-0 sm:block"
      ></div>
      <div
        className={`absolute ${position} z-[21] inline-flex min-w-[16rem] max-w-[21rem] rounded-lg bg-white p-4 text-sm shadow-right`}
      >
        <div className="flex w-full flex-col">
          {/* Filter breadcrumbs */}
          <InventoryFilterBreadcrumbs step={step} goTo={goTo} />
          <div className="mt-2"></div>

          {/* Filter summary */}
          {step !== 0 && data && data.field && (
            <InventoryFilterSummary data={data} resetData={resetData} />
          )}
          <div className="mt-2"></div>

          {/* Filter steps - 1/3 filter field */}
          {step === 0 && <InventoryFilterField handleField={handleField} />}

          {/* Filter steps - 2/3 filter operator */}
          {step === 1 && (
            <InventoryFilterOperator
              data={data}
              handleTagKey={handleTagKey}
              handleOperator={handleOperator}
            />
          )}

          {/* Filter steps - 3/3 filter value */}
          {step === 2 && (
            <form
              onSubmit={e => {
                e.preventDefault();
                filter();
                if (closeDropdownAfterAdd) toggle();
              }}
            >
              <div className="-mr-4 max-h-[calc(100vh-23rem)] overflow-auto pb-2">
                <InventoryFilterValue
                  data={data}
                  handleValueCheck={handleValueCheck}
                  handleValueInput={handleValueInput}
                  cleanValues={cleanValues}
                  showToast={showToast}
                  costBetween={costBetween}
                  handleCostBetween={handleCostBetween}
                />
              </div>
              {inlineError.hasError && (
                <p className="pb-4 text-xs font-medium text-red-500">
                  {inlineError.message}
                </p>
              )}
              <div className="-mx-4 -mb-4 flex justify-end border-t px-4 py-4">
                <Button
                  type="submit"
                  disabled={
                    data &&
                    data.operator !== 'IS_EMPTY' &&
                    data.operator !== 'IS_NOT_EMPTY' &&
                    data.operator !== 'EXISTS' &&
                    data.operator !== 'NOT_EXISTS' &&
                    data.values &&
                    data.values.length === 0
                  }
                >
                  Apply filter
                </Button>
              </div>
            </form>
          )}
        </div>
      </div>
    </>
  );
}
