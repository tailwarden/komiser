import Button from '../../../button/Button';
import { InventoryFilterDataProps } from '../../hooks/useInventory';
import inventoryFilterFieldOptions from './InventoryFilterFieldOptions';

type InventoryFilterSummaryProps = {
  id?: number;
  bg?: 'white';
  data: InventoryFilterDataProps;
  deleteFilter?: (idx: number) => void;
  resetData?: () => void;
};

function InventoryFilterSummary({
  id,
  bg,
  data,
  deleteFilter,
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
    if (param === 'IS_EMPTY' && data.field !== 'tags') return 'is empty';
    if (param === 'IS_EMPTY' && data.field === 'tags') return 'which are empty';
    if (param === 'IS_NOT_EMPTY' && data.field !== 'tags')
      return 'is not empty';
    if (param === 'IS_NOT_EMPTY' && data.field === 'tags')
      return 'which are not empty';
    if (param === 'EQUAL') return 'is equal to';
    if (param === 'BETWEEN') return 'is between';
    if (param === 'GREATER_THAN') return 'is greater than';
    if (param === 'LESS_THAN') return 'is less than';
    return param;
  }

  return (
    <div
      className={`${bg ? 'bg-white' : 'bg-black-100'} ${
        deleteFilter || resetData ? 'pr-12' : 'pr-4'
      } relative flex text-black-900/70 p-2 pr-12 text-xs rounded max-w-[calc(100vw-250px)] md:max-w-[calc(100vw-400px)] overflow-hidden`}
    >
      {(deleteFilter || resetData) && (
        <div
          className={`absolute bottom-[.35rem] right-1 ${
            bg ? 'bg-white' : 'bg-black-100'
          }`}
        >
          <Button
            size="xs"
            style="ghost"
            onClick={() => {
              if (deleteFilter) {
                deleteFilter(id!);
              } else {
                resetData!();
              }
            }}
          >
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
      )}

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
              <span>
                {data.field === 'cost' && '$'}
                {value}
              </span>
              {data.values.length > 1 && idx < data.values.length - 1 && (
                <span className="ml-1 font-medium text-black-900">
                  {data.field === 'cost' ? 'and' : 'or'}
                </span>
              )}
            </p>
          ))}
      </div>
    </div>
  );
}

export default InventoryFilterSummary;
