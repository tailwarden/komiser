import regex from '../../../../utils/regex';
import Button from '../../../button/Button';
import Input from '../../../input/Input';
import { InventoryFilterDataProps } from '../../hooks/useInventory';

type InventoryFilterOperatorProps = {
  data: InventoryFilterDataProps;
  handleOperator: (operator: InventoryFilterDataProps['operator']) => void;
  handleTagKey: (newValue: { tagKey: string }) => void;
};

export type InventoryFilterOperatorOptionsProps = {
  label: string;
  value: InventoryFilterDataProps['operator'];
};

const inventoryFilterOperatorOptions: InventoryFilterOperatorOptionsProps[] = [
  { label: 'is', value: 'IS' },
  { label: 'is not', value: 'IS_NOT' },
  { label: 'contains', value: 'CONTAINS' },
  { label: 'does not contain', value: 'NOT_CONTAINS' },
  { label: 'is empty', value: 'IS_EMPTY' },
  { label: 'is not empty', value: 'IS_NOT_EMPTY' }
];

const inventoryFilterOperatorAllTagsOptions: InventoryFilterOperatorOptionsProps[] =
  [
    { label: 'which are empty', value: 'IS_EMPTY' },
    { label: 'which are not empty', value: 'IS_NOT_EMPTY' }
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
        <div className="pl-1 pt-2 pb-2">
          <Input
            type="text"
            name="tagKey"
            label="Tag key"
            value={data.tagKey}
            regex={regex.required}
            error="Please provide a tag key"
            action={handleTagKey}
            autofocus={true}
          />
        </div>
      )}

      {/* Operators list */}
      {data.field !== 'tags' &&
        inventoryFilterOperatorOptions.map((option, idx) => (
          <Button
            key={idx}
            size="sm"
            style="ghost"
            align="left"
            gap="md"
            disabled={data.field === 'tag' && !data.tagKey}
            transition={false}
            onClick={() => handleOperator(option.value)}
          >
            {option.label}
          </Button>
        ))}
      {data.field === 'tags' &&
        inventoryFilterOperatorAllTagsOptions.map((option, idx) => (
          <Button
            key={idx}
            size="sm"
            style="ghost"
            align="left"
            gap="md"
            transition={false}
            onClick={() => handleOperator(option.value)}
          >
            {option.label}
          </Button>
        ))}
    </div>
  );
}

export default InventoryFilterOperator;
