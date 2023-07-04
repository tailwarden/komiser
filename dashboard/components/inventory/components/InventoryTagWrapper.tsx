import Input from '../../input/Input';
import { Tag } from '../hooks/useInventory/types/useInventoryTypes';

type InventoryTagWrapperProps = {
  tag: Tag;
  id: number;
  handleChange: (newData: Partial<Tag>, id?: number) => void;
};

function InventoryTagWrapper({
  tag,
  id,
  handleChange
}: InventoryTagWrapperProps) {
  return (
    <div className="grid flex-grow grid-cols-2 gap-4">
      <Input
        id={id}
        name="key"
        label="Key"
        type="text"
        error="Please provide a valid key name."
        action={handleChange}
        value={tag.key}
        maxLength={64}
      />
      <Input
        id={id}
        name="value"
        label="Value"
        type="text"
        error="Please provide a valid value."
        action={handleChange}
        value={tag.value}
        maxLength={64}
      />
    </div>
  );
}

export default InventoryTagWrapper;
