import Button from '../../../button/Button';
import { InventoryFilterDataProps } from '../../hooks/useFilterWizard';
import inventoryFilterOptions from './InventoryFilterFieldOptions';

type InventoryFilterSummaryProps = {
  data: InventoryFilterDataProps;
  resetData: () => void;
};

function InventoryFilterSummary({
  data,
  resetData
}: InventoryFilterSummaryProps) {
  const index = inventoryFilterOptions.findIndex(
    option => option.value === data.field
  );

  function getData(param: 'icon' | 'label') {
    if (param === 'icon') return inventoryFilterOptions[index].icon;
    if (param === 'label') return inventoryFilterOptions[index].label;
    return param;
  }

  return (
    <div className="flex justify-between gap-4 bg-black-100 p-2 text-xs rounded mb-2">
      <div className="flex items-center gap-1">
        <div className="scale-75">{getData('icon')}</div>
        <p>{getData('label')}</p>
        <p>{data.operator}</p>
        <p>{data.values?.map(value => value)}</p>
      </div>
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
  );
}

export default InventoryFilterSummary;
