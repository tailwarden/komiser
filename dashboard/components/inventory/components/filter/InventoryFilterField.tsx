import Button from '../../../button/Button';
import inventoryFilterOptions from './InventoryFilterFieldOptions';

type InventoryFilterFieldProps = {
  handleField: (field: string) => void;
};

function InventoryFilterField({ handleField }: InventoryFilterFieldProps) {
  return (
    <>
      {inventoryFilterOptions.map((option, idx) => (
        <Button
          key={idx}
          size="sm"
          style="ghost"
          align="left"
          gap="md"
          onClick={() => handleField(option.value)}
        >
          {option.icon}
          {option.label}
        </Button>
      ))}
    </>
  );
}

export default InventoryFilterField;
