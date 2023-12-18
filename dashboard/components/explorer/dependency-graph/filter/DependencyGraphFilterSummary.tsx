import { InventoryFilterData } from '@components/inventory/hooks/useInventory/types/useInventoryTypes';
import Button from '@components/button/Button';
import CloseIcon from '@components/icons/CloseIcon';
import DependencyGraphFilterOptions from './DependencyGraphFilterOptions';

type DependencyGraphFilterSummaryProps = {
  id?: number;
  bg?: 'white';
  data: InventoryFilterData;
  deleteFilter?: (idx: number) => void;
  resetData?: () => void;
};

function DependencyGraphFilterSummary({
  id,
  bg,
  data,
  deleteFilter,
  resetData
}: DependencyGraphFilterSummaryProps) {
  const index = DependencyGraphFilterOptions.findIndex(
    option => option.value === data.field
  );

  function getField(param: 'icon' | 'label') {
    if (param === 'icon') return DependencyGraphFilterOptions[index].icon;
    if (param === 'label') return DependencyGraphFilterOptions[index].label;
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
      className={`${
        bg ? 'bg-white' : 'bg-background-ds'
      }  relative flex h-6 w-fit max-w-[calc(100vw-250px)] items-center gap-1 overflow-hidden rounded px-2 text-xs text-gray-950 md:max-w-[calc(100vw-400px)]`}
    >
      <div className="flex items-center gap-1 whitespace-nowrap">
        <div className="scale-75">{getField('icon')}</div>
        <p>{getField('label')}</p>
        {data.tagKey && <p>: {data.tagKey}</p>}
        {data.operator && (
          <>
            <span>:</span>
            <span className="">{getOperator(data.operator)}</span>
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
                <span className="ml-1 ">
                  {data.field === 'cost' && data.operator === 'BETWEEN'
                    ? 'and'
                    : 'or'}
                </span>
              )}
            </p>
          ))}
      </div>
      {(deleteFilter || resetData) && (
        <div>
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
            <CloseIcon width={16} height={20} />
          </Button>
        </div>
      )}
    </div>
  );
}

export default DependencyGraphFilterSummary;
