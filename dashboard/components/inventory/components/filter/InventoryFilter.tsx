import Button from '../../../button/Button';
import useFilterWizard from '../../hooks/useFilterWizard';
import InventoryFilterField from './InventoryFilterField';
import InventoryFilterOperator from './InventoryFilterOperator';
import InventoryFilterSummary from './InventoryFilterSummary';
import InventoryFilterValue from './InventoryFilterValue';

function InventoryFilter() {
  const {
    toggle,
    isOpen,
    step,
    goTo,
    handleField,
    handleOperator,
    data,
    resetData
  } = useFilterWizard();

  return (
    <div className="relative">
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
          <div
            onClick={toggle}
            className="hidden sm:block fixed inset-0 z-20 bg-transparent opacity-0 animate-fade-in"
          ></div>
          <div className="absolute flex flex-col min-w-[16rem] left-0 top-12 bg-white p-4 shadow-xl text-sm rounded-lg z-[21]">
            <div className="flex gap-2 text-black-400">
              <p
                className="cursor-pointer hover:text-black-900"
                onClick={() => goTo(0)}
              >
                Fields
              </p>
              {step !== 0 && (
                <>
                  <p>&gt;</p>
                  <p
                    className="cursor-pointer hover:text-black-900"
                    onClick={() => goTo(1)}
                  >
                    Operator
                  </p>
                  {step === 2 && (
                    <>
                      <p>&gt;</p>
                      <p
                        className="cursor-pointer hover:text-black-900"
                        onClick={() => goTo(2)}
                      >
                        Value
                      </p>
                    </>
                  )}
                </>
              )}
            </div>
            <div className="mt-2"></div>
            {step !== 0 && data && data.field && (
              <InventoryFilterSummary data={data} resetData={resetData} />
            )}
            {step === 0 && <InventoryFilterField handleField={handleField} />}
            {step === 1 && (
              <InventoryFilterOperator handleOperator={handleOperator} />
            )}
            {step === 2 && (
              <>
                <InventoryFilterValue />
                <div className="mt-2"></div>
                <Button>Apply filter</Button>
              </>
            )}
          </div>
        </>
      )}
    </div>
  );
}

export default InventoryFilter;
