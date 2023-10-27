import type { Meta, StoryObj } from '@storybook/react';

import { useEffect, useState } from 'react';
import Upload, { UploadProps } from './Upload';

function UploadWrapper({
  multiple,
  fileOrFiles,
  handleChange,
  onClose,
  ...otherProps
}: UploadProps) {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [selectedFiles, setSelectedFiles] = useState<File[] | null>(null);

  useEffect(() => {
    setSelectedFile(null);
    setSelectedFiles(null);
  }, [multiple]);

  const uploadFile = (file: File | File[] | null): void => {
    if (file instanceof FileList) {
      const filesArray = Array.from(file);
      setSelectedFiles(filesArray);
    } else if (file instanceof File) {
      setSelectedFile(file);
    } else {
      setSelectedFile(null);
      setSelectedFiles(null);
    }
  };

  return (
    // it's impossible to define a true/false type in storybook
    // so we ignore the next type error because true|false != boolean for some reason \o/
    // @ts-ignore
    multiple ? (
      <Upload
        multiple={multiple}
        fileOrFiles={selectedFiles}
        handleChange={uploadFile}
        onClose={() => setSelectedFiles(null)}
        {...otherProps}
      />
    ) : (
      <Upload
        multiple={multiple}
        fileOrFiles={selectedFile}
        handleChange={uploadFile}
        onClose={() => setSelectedFile(null)}
        {...otherProps}
      />
    )
  );
}

const meta: Meta<typeof Upload> = {
  title: 'Komiser/Upload',
  component: UploadWrapper,
  decorators: [
    Story => (
      <div className="h-full w-full rounded bg-white p-0.5">{Story()}</div>
    )
  ],
  tags: ['autodocs'],
  argTypes: {
    name: {
      control: 'text',
      description: 'the name for your form (if exist)',
      defaultValue: 'attachment'
    },
    multiple: {
      control: 'boolean',
      description:
        'a boolean to determine whether the multiple files is enabled or not',
      defaultValue: false
    },
    disabled: {
      control: 'boolean',
      description: 'disables the input',
      defaultValue: false
    },
    required: {
      control: 'boolean',
      description: 'Conditionally set the input field as required',
      defaultValue: false
    },
    hoverTitle: {
      control: 'text',
      description: 'text appears(hover) when trying to drop a file',
      defaultValue: 'drop here'
    },
    maxSize: {
      control: 'number',
      description: 'the maximum size of the file (number in mb)',
      defaultValue: 37
    },
    minSize: {
      control: 'number',
      description: 'the minimum size of the file (number in mb)',
      defaultValue: 0
    }
  }
};

export default meta;

type Story = StoryObj<typeof Upload>;

export const SingleFile: Story = {
  args: {
    name: 'attachment',
    multiple: false,
    disabled: false,
    hoverTitle: 'drop here',
    maxSize: 37,
    minSize: 0
  }
};

export const MultipleFiles: Story = {
  args: {
    name: 'attachment',
    multiple: true,
    disabled: false,
    hoverTitle: 'drop here',
    maxSize: 37,
    minSize: 0
  }
};
