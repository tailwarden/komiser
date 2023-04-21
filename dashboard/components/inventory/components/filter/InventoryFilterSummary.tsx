import Button from '../../../button/Button';
import CloseIcon from '../../../icons/CloseIcon';
import { InventoryFilterData } from '../../hooks/useInventory/types/useInventoryTypes';
import inventoryFilterFieldOptions from './InventoryFilterFieldOptions';

type InventoryFilterSummaryProps = {
  id?: number;
  bg?: 'white';
  data: InventoryFilterData;
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

  function getOperator(param: InventoryFilterData['operator']) {
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
    if (param === 'EXISTS') return 'does exist';
    if (param === 'NOT_EXISTS') return 'does not exist';
    return param;
  }

  return (
    <div
      className={`${bg ? 'bg-white' : 'bg-black-100'} ${
        deleteFilter || resetData ? 'pr-12' : 'pr-4'
      } relative flex max-w-[calc(100vw-250px)] overflow-hidden rounded p-2 pr-12 text-xs text-black-900/70 md:max-w-[calc(100vw-400px)]`}
    >
      {(deleteFilter || resetData) && (
        <div
          className={`absolute bottom-[.25rem] ${
            deleteFilter ? 'right-1.5' : 'right-0'
          } ${bg ? 'bg-white' : 'bg-black-100'}`}
        >
          <Button
            size="xxs"
            style="ghost"
            onClick={() => {
              if (deleteFilter) {
                deleteFilter(id!);
              } else {
                resetData!();
              }
            }}
          >
            <CloseIcon width={20} height={24} />
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
                  {data.field === 'cost' && data.operator === 'BETWEEN'
                    ? 'and'
                    : 'or'}
                </span>
              )}
            </p>
          ))}
      </div>
    </div>
  );
}

export default InventoryFilterSummary;
