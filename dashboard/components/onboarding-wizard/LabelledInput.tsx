interface LabelledInputProps {
  label: string;
  subLabel?: string;
  placeholder: string;
}

function LabelledInput({
  label,
  subLabel,
  placeholder = ''
}: LabelledInputProps) {
  return (
    <div>
      <label htmlFor="input-group-1" className="mb-2 block text-gray-700">
        {label}
      </label>

      {subLabel && (
        <span className="-mt-[5px] mb-2 block text-xs leading-4 text-black-400">
          {subLabel}
        </span>
      )}

      <div className="relative mb-6">
        <input
          type="text"
          id="input-group-1"
          className="block w-full rounded py-4 pl-3 text-sm text-black-900 outline outline-black-200 focus:outline-2 focus:outline-primary"
          placeholder={placeholder}
        />
      </div>
    </div>
  );
}

export default LabelledInput;
