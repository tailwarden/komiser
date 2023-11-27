type InventoryFilterBreadcrumbsProps = {
  step: number;
  goTo: (index: number) => void;
};

function InventoryFilterBreadcrumbs({
  step,
  goTo
}: InventoryFilterBreadcrumbsProps) {
  return (
    <div className="flex gap-2 text-xs text-gray-500">
      <p
        className={`cursor-pointer ${
          step === 0
            ? 'text-gray-950 hover:text-gray-950'
            : 'hover:text-gray-700'
        }`}
        onClick={() => goTo(0)}
      >
        Fields
      </p>
      {step !== 0 && (
        <>
          <p>&gt;</p>
          <p
            className={`cursor-pointer ${
              step === 1
                ? 'text-gray-950 hover:text-gray-950'
                : 'hover:text-gray-700'
            }`}
            onClick={() => goTo(1)}
          >
            Operator
          </p>
          {step === 2 && (
            <>
              <p>&gt;</p>
              <p
                className={`cursor-pointer ${
                  step === 2
                    ? 'text-gray-950 hover:text-gray-950'
                    : 'hover:text-gray-700'
                }`}
                onClick={() => goTo(2)}
              >
                Value
              </p>
            </>
          )}
        </>
      )}
    </div>
  );
}

export default InventoryFilterBreadcrumbs;
