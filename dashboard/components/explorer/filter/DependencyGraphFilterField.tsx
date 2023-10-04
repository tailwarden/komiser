import Button from '@components/button/Button';
import DependencyGraphFilterFieldOptions from './DependencyGraphFilterOptions';

type DependencyGraphFilterFieldProps = {
  handleField: (field: string) => void;
};

function DependencyGraphFilterField({
  handleField
}: DependencyGraphFilterFieldProps) {
  return (
    <>
      {DependencyGraphFilterFieldOptions.map((option, idx) => (
        <Button
          key={idx}
          size="sm"
          style="dropdown"
          gap="md"
          transition={false}
          onClick={() => handleField(option.value)}
        >
          {option.icon}
          {option.label}
        </Button>
      ))}
    </>
  );
}

export default DependencyGraphFilterField;
