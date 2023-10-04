import Button from '../../../button/Button';
import Input from '../../../input/Input';
import { InventoryFilterData } from '../../hooks/useInventory/types/useInventoryTypes';

type InventoryFilterOperatorProps = {
  data: InventoryFilterData;
  handleOperator: (operator: InventoryFilterData['operator']) => void;
  handleTagKey: (newValue: { tagKey: string }) => void;
};

export type InventoryFilterOperatorOptionsProps = {
  label: string;
  value: InventoryFilterData['operator'];
};

const inventoryFilterOperatorOptions: InventoryFilterOperatorOptionsProps[] = [
  { label: 'is', value: 'IS' },
  { label: 'is not', value: 'IS_NOT' },
  { label: 'contains', value: 'CONTAINS' },
  { label: 'does not contain', value: 'NOT_CONTAINS' },
  { label: 'is empty', value: 'IS_EMPTY' },
  { label: 'is not empty', value: 'IS_NOT_EMPTY' }
];

const inventoryFilterOperatorSpecificTagsOptions: InventoryFilterOperatorOptionsProps[] =
  [
    { label: 'does exist', value: 'EXISTS' },
    { label: 'does not exist', value: 'NOT_EXISTS' }
  ];

const inventoryFilterOperatorAllTagsOptions: InventoryFilterOperatorOptionsProps[] =
  [
    { label: 'which are empty', value: 'IS_EMPTY' },
    { label: 'which are not empty', value: 'IS_NOT_EMPTY' }
  ];

const inventoryFilterOperatorCostOptions: InventoryFilterOperatorOptionsProps[] =
  [
    {
      label: 'is equal to',
      value: 'EQUAL'
    },
    {
      label: 'is between',
      value: 'BETWEEN'
    },
    {
      label: 'is greater than',
      value: 'GREATER_THAN'
    },
    {
      label: 'is less than',
      value: 'LESS_THAN'
    }
  ];

const inventoryFilterRelationsOptions: InventoryFilterOperatorOptionsProps[] = [
  {
    label: 'is equal to',
    value: 'EQUAL'
  },
  {
    label: 'is greater than',
    value: 'GREATER_THAN'
  },
  {
    label: 'is less than',
    value: 'LESS_THAN'
  }
];

function InventoryFilterOperator({
  data,
  handleOperator,
  handleTagKey
}: InventoryFilterOperatorProps) {
  return (
    <div className="flex flex-col">
      {/* If field is tag, ask for tag key */}
      {data.field === 'tag' && (
        <div className="pb-2 pl-1 pt-2">
          <Input
            type="text"
            name="tagKey"
            label="Tag key"
            value={data.tagKey}
            error="Please provide a tag key"
            action={handleTagKey}
            autofocus={true}
            maxLength={64}
          />
        </div>
      )}

      {/* Operators list which are not tags or cost */}
      {data.field !== 'tags' &&
        data.field !== 'cost' &&
        data.field !== 'relations' &&
        inventoryFilterOperatorOptions.map((option, idx) => (
          <Button
            key={idx}
            style="dropdown"
            disabled={data.field === 'tag' && !data.tagKey}
            onClick={() => handleOperator(option.value)}
          >
            {option.label}
          </Button>
        ))}

      {/* Operators list for specific tags */}
      {data.field === 'tag' &&
        inventoryFilterOperatorSpecificTagsOptions.map((option, idx) => (
          <Button
            key={idx}
            style="dropdown"
            disabled={data.field === 'tag' && !data.tagKey}
            onClick={() => handleOperator(option.value)}
          >
            {option.label}
          </Button>
        ))}

      {/* Operators list for tags */}
      {data.field === 'tags' &&
        inventoryFilterOperatorAllTagsOptions.map((option, idx) => (
          <Button
            key={idx}
            style="dropdown"
            onClick={() => handleOperator(option.value)}
          >
            {option.label}
          </Button>
        ))}

      {/* Operators list for cost */}
      {data.field === 'cost' &&
        inventoryFilterOperatorCostOptions.map((option, idx) => (
          <Button
            key={idx}
            style="dropdown"
            onClick={() => handleOperator(option.value)}
          >
            {option.label}
          </Button>
        ))}

      {/* Operators list for cost */}
      {data.field === 'relations' &&
        inventoryFilterRelationsOptions.map((option, idx) => (
          <Button
            key={option.value}
            style="dropdown"
            onClick={() => handleOperator(option.value)}
          >
            {option.label}
          </Button>
        ))}
    </div>
  );
}

export default InventoryFilterOperator;
