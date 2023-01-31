import regex from '../../../utils/regex';
import Input from '../../input/Input';
import { Tag } from '../hooks/useInventory';

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
    <div className="grid flex-grow grid-cols-2 gap-6">
      <Input
        id={id}
        name="key"
        label="Key"
        type="text"
        regex={regex.required}
        error="Please provide a valid key name."
        action={handleChange}
        value={tag.key}
      />
      <Input
        id={id}
        name="value"
        label="Value"
        type="text"
        regex={regex.required}
        error="Please provide a valid value."
        action={handleChange}
        value={tag.value}
      />
    </div>
  );
}

export default InventoryTagWrapper;
