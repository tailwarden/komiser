import regex from '../../../utils/regex';
import Input from '../../input/Input';
import { Tag } from '../hooks/useInventory';

type TagWrapperProps = {
  tag: Tag;
  id: number;
  handleChange: (newData: Partial<Tag>, id?: number) => void;
};

function TagWrapper({ tag, id, handleChange }: TagWrapperProps) {
  return (
    <div className="flex-grow grid grid-cols-2 gap-6">
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

export default TagWrapper;
