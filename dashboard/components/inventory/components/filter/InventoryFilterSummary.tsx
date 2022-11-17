import Button from '../../../button/Button';
import { InventoryFilterDataProps } from '../../hooks/useInventory';
import inventoryFilterFieldOptions from './InventoryFilterFieldOptions';

type InventoryFilterSummaryProps = {
  bg?: 'white';
  data: InventoryFilterDataProps;
  resetData: () => void;
};

function InventoryFilterSummary({
  bg,
  data,
  resetData
}: InventoryFilterSummaryProps) {
  const index = inventoryFilterFieldOptions.findIndex(
    option => option.value === data.field
  );

  function getField(param: 'icon' | 'label') {
    if (param === 'icon') return inventoryFilterFieldOptions[index].icon;
    if (param === 'label') return inventoryFilterFieldOptions[index].label;
    return param;
  }

  function getOperator(param: InventoryFilterDataProps['operator']) {
    if (param === 'IS') return 'is';
    if (param === 'IS_NOT') return 'is not';
    if (param === 'CONTAINS') return 'contains';
    if (param === 'NOT_CONTAINS') return 'does not contain';
    if (param === 'IS_EMPTY') return 'is empty';
    if (param === 'IS_NOT_EMPTY') return 'is not empty';
    return param;
  }

  return (
    <div
      className={`${
        bg ? 'bg-white' : 'bg-black-100'
      } relative flex text-black-900/70 p-2 pr-12 text-xs rounded max-w-[calc(100vw-250px)] md:max-w-[calc(100vw-400px)] overflow-hidden`}
    >
      <div
        className={`absolute bottom-[.35rem] right-1 ${
          bg ? 'bg-white' : 'bg-black-100'
        }`}
      >
        <Button size="xs" style="ghost" onClick={resetData}>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            fill="none"
            viewBox="0 0 24 24"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M7.757 16.243l8.486-8.486M16.243 16.243L7.757 7.757"
            ></path>
          </svg>
        </Button>
      </div>
      <div className="flex items-center gap-1 whitespace-nowrap">
        <div className="scale-75">{getField('icon')}</div>
        <p>{getField('label')}</p>
        {data.tagKey && <p>: {data.tagKey}</p>}
        {data.operator && (
          <>
            <span>:</span>
            <span className="font-medium text-black-900">
              {getOperator(data.operator)}
            </span>
          </>
        )}
        {data.values &&
          data.values.length > 0 &&
          data.values.map((value, idx) => (
            <p key={idx}>
              {idx === 0 && <span className="mr-1">:</span>}
              <span>{value}</span>
              {data.values.length > 1 && idx < data.values.length - 1 && (
                <span className="ml-1 font-medium text-black-900">or</span>
              )}
            </p>
          ))}
      </div>
    </div>
  );
}

export default InventoryFilterSummary;
