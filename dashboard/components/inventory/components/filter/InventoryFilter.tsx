import { NextRouter } from 'next/router';
import Button from '../../../button/Button';
import { ToastProps } from '../../../toast/hooks/useToast';
import useFilterWizard from '../../hooks/useFilterWizard';
import { InventoryFilterDataProps } from '../../hooks/useInventory';
import InventoryFilterBreadcrumbs from './InventoryFilterBreadcrumbs';
import InventoryFilterField from './InventoryFilterField';
import InventoryFilterOperator from './InventoryFilterOperator';
import InventoryFilterSummary from './InventoryFilterSummary';
import InventoryFilterValue from './InventoryFilterValue';

type InventoryFilterProps = {
  router: NextRouter;
  displayedFilters: InventoryFilterDataProps[] | undefined;
  setSkippedSearch: (number: number) => void;
  setToast: (toast: ToastProps | undefined) => void;
  deleteFilter: (idx: number) => void;
};

function InventoryFilter({
  router,
  displayedFilters,
  setSkippedSearch,
  setToast,
  deleteFilter
}: InventoryFilterProps) {
  const {
    toggle,
    isOpen,
    step,
    goTo,
    handleField,
    handleOperator,
    handleTagKey,
    handleValueCheck,
    handleValueInput,
    data,
    resetData,
    cleanValues,
    filter
  } = useFilterWizard({ router, setSkippedSearch });

  return (
    <div>
      {/* Active filters list */}
      {Object.keys(router.query).length > 0 &&
        displayedFilters &&
        displayedFilters.length > 0 && (
          <div className="flex items-center gap-4 bg-white py-2 px-6 rounded-lg absolute left-0 right-0 top-[64px]">
            <span className="text-sm text-black-400">Filters:</span>
            {displayedFilters.map((activeFilter, idx) => (
              <InventoryFilterSummary
                key={idx}
                id={idx}
                data={activeFilter}
                deleteFilter={deleteFilter}
              />
            ))}
          </div>
        )}

      {/* Dropdown button toggle */}
      <button
        className={`flex items-center font-medium text-sm rounded-lg h-[2.5rem] px-3 gap-2 text-black-900/60 ${
          isOpen
            ? 'bg-black-400/10'
            : 'bg-transparent hover:bg-black-400/10 active:bg-black-400/20'
        }`}
        onClick={toggle}
      >
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
      </button>

      {/* Dropdown */}
      {isOpen && (
        <>
          {/* Dropdown transparent backdrop */}
          <div
            onClick={toggle}
            className="hidden sm:block fixed inset-0 z-20 bg-transparent opacity-0 animate-fade-in"
          ></div>
          <div className="absolute inline-flex min-w-[16rem] right-0 top-12 bg-white p-4 shadow-xl text-sm rounded-lg z-[21]">
            {/* <div className="absolute inline-flex min-w-[16rem] left-0 top-12 bg-white p-4 shadow-xl text-sm rounded-lg z-[21]"> */}
            <div className="flex flex-col w-full">
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
                  }}
                >
                  <div className="max-h-[calc(100vh-23rem)] overflow-auto -mr-4 pb-2">
                    <InventoryFilterValue
                      data={data}
                      handleValueCheck={handleValueCheck}
                      handleValueInput={handleValueInput}
                      cleanValues={cleanValues}
                      setToast={setToast}
                    />
                  </div>
                  <div className="border-t -mx-4 -mb-4 px-4 py-4">
                    <Button
                      type="submit"
                      disabled={
                        data &&
                        data.operator !== 'IS_EMPTY' &&
                        data.operator !== 'IS_NOT_EMPTY' &&
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
      )}
    </div>
  );
}

export default InventoryFilter;
