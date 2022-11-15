import Button from '../../../button/Button';
import { InventoryFilterDataProps } from '../../hooks/useFilterWizard';

type InventoryFilterOperatorProps = {
  handleOperator: (operator: InventoryFilterDataProps['operator']) => void;
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

function InventoryFilterOperator({
  handleOperator
}: InventoryFilterOperatorProps) {
  return (
    <div className="flex flex-col">
      {inventoryFilterOperatorOptions.map((option, idx) => (
        <Button
          key={idx}
          size="sm"
          style="ghost"
          align="left"
          gap="md"
          onClick={() => handleOperator(option.value)}
        >
          {option.label}
        </Button>
      ))}
    </div>
  );
}

export default InventoryFilterOperator;
