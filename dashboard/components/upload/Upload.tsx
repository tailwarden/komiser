import React from 'react';
import { FileUploader } from 'react-drag-drop-files';

type BaseUploadProps = {
  name?: string;
  label?: string;
  required?: boolean;
  disabled?: boolean;
  hoverTitle?: string;
  classes?: string;
  childClassName?: string;
  types?: string[];
  onTypeError?: (error: string) => void;
  children?: any;
  maxSize?: number;
  minSize?: number;
  onSizeError?: (error: string) => void;
  onDrop?: (file: File | null) => void;
  onSelect?: (file: File | null) => void;
  onClose: () => void;
  onDraggingStateChange?: () => void;
  dropMessageStyle?: React.CSSProperties;
};

export type SingleUploadProps = BaseUploadProps & {
  multiple: false;
  fileOrFiles: File | null;
  handleChange: (file: File | null) => void;
};

export type MultipleUploadProps = BaseUploadProps & {
  multiple: true;
  fileOrFiles: File[] | null;
  handleChange: (files: File[] | null) => void;
};

export type UploadProps = SingleUploadProps | MultipleUploadProps;

const FILE_TYPES = ['JPG', 'PNG', 'GIF', 'TXT', 'LOG', 'MP4', 'AVI', 'MOV'];
const MAX_FILE_SIZE_MB = 37;

const defaultDropMessageStyle: React.CSSProperties = {
  width: '100%',
  height: '100%',
  position: 'absolute',
  background: '#F4F9F9',
  top: 0,
  right: 2,
  display: 'flex',
  flexGrow: 2,
  opacity: 1,
  zIndex: 20,
  color: '#008484',
  fontSize: 14,
  border: 'none'
};

function Upload({
  name = 'attachment',
  multiple = false,
  label,
  required,
  disabled,
  hoverTitle = 'drop here',
  fileOrFiles,
  handleChange,
  classes,
  childClassName,
  types = FILE_TYPES,
  onTypeError,
  maxSize = MAX_FILE_SIZE_MB,
  minSize,
  onSizeError,
  onDrop,
  onSelect,
  onClose,
  onDraggingStateChange,
  dropMessageStyle = defaultDropMessageStyle
}: UploadProps) {
  const defaultChildClassName =
    fileOrFiles === null
      ? `grow bg-white cursor-pointer rounded border-2 border-dashed border-gray-200 py-5 text-center text-xs transition hover:border-[#B6EAEA] hover:bg-gray-50 w-full`
      : `grow bg-white min-h-full rounded border-2 border-[#B6EAEA] text-center text-xs transition w-full`;

  return (
    <FileUploader
      name={name}
      multiple={multiple}
      label={label}
      required={required}
      disabled={disabled || fileOrFiles !== null}
      fileOrFiles={fileOrFiles}
      handleChange={handleChange}
      classes={classes}
      types={types}
      hoverTitle={hoverTitle}
      onTypeError={onTypeError}
      maxSize={maxSize}
      minSize={minSize}
      onSizeError={onSizeError}
      onDrop={onDrop}
      onSelect={onSelect}
      onDraggingStateChange={onDraggingStateChange}
      dropMessageStyle={dropMessageStyle}
    >
      <div className={childClassName || defaultChildClassName}>
        {fileOrFiles === null && (
          <div className="">
            <svg
              className="mb-2 inline-block"
              width="25"
              height="24"
              viewBox="0 0 25 24"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M9.5 17V11L7.5 13"
                stroke="#0C1717"
                strokeWidth="1.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
              <path
                d="M9.5 11L11.5 13"
                stroke="#0C1717"
                strokeWidth="1.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
              <path
                d="M22.5 10V15C22.5 20 20.5 22 15.5 22H9.5C4.5 22 2.5 20 2.5 15V9C2.5 4 4.5 2 9.5 2H14.5"
                stroke="#0C1717"
                strokeWidth="1.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
              <path
                d="M22.5 10H18.5C15.5 10 14.5 9 14.5 6V2L22.5 10Z"
                stroke="#0C1717"
                strokeWidth="1.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
            <p>
              Drag and drop or{' '}
              <span className="cursor-pointer text-darkcyan-500">
                choose a file
              </span>
            </p>
          </div>
        )}
        {fileOrFiles !== null && (
          <div className="flex h-full flex-col gap-2">
            {multiple ? (
              (fileOrFiles as File[])?.map((file: File, index: number) => (
                <div className="relative h-full w-full px-6 py-5" key={index}>
                  {/* Render the multiple files */}
                  <div className="flex h-full w-full items-center justify-items-center gap-2">
                    <svg
                      width="40"
                      height="40"
                      viewBox="0 0 40 40"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        d="M35 11.6666V28.3333C35 33.3333 32.5 36.6666 26.6667 36.6666H13.3333C7.5 36.6666 5 33.3333 5 28.3333V11.6666C5 6.66659 7.5 3.33325 13.3333 3.33325H26.6667C32.5 3.33325 35 6.66659 35 11.6666Z"
                        stroke="#0C1717"
                        strokeWidth="1.5"
                        strokeMiterlimit="10"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      />
                      <path
                        d="M24.167 7.5V10.8333C24.167 12.6667 25.667 14.1667 27.5003 14.1667H30.8337"
                        stroke="#0C1717"
                        strokeWidth="1.5"
                        strokeMiterlimit="10"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      />
                      <path
                        d="M13.333 21.6667H19.9997"
                        stroke="#0C1717"
                        strokeWidth="1.5"
                        strokeMiterlimit="10"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      />
                      <path
                        d="M13.333 28.3333H26.6663"
                        stroke="#0C1717"
                        strokeWidth="1.5"
                        strokeMiterlimit="10"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      />
                    </svg>
                    <div className="flex-1 text-left">
                      <p>{file.name}</p>
                      <p className="text-gray-700">
                        {(file.size / (1024 * 1024)).toFixed(2)}
                        MB
                      </p>
                    </div>
                  </div>
                  <a
                    onClick={onClose}
                    className={`absolute right-4 top-4 block h-4 w-4 cursor-pointer ${
                      index !== 0 && 'hidden'
                    }`}
                    aria-disabled={!file}
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="16"
                      height="17"
                      viewBox="0 0 16 17"
                      fill="none"
                    >
                      <path
                        d="M4.66602 12.079L11.3327 5.41235"
                        stroke="#0C1717"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      />
                      <path
                        d="M11.3327 12.079L4.66602 5.41235"
                        stroke="#0C1717"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      />
                    </svg>
                  </a>
                </div>
              ))
            ) : (
              <div className="relative h-full w-full px-6 py-5">
                {/* Render the single file */}
                <div className="flex h-full w-full items-center justify-items-center gap-2">
                  <svg
                    width="40"
                    height="40"
                    viewBox="0 0 40 40"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      d="M35 11.6666V28.3333C35 33.3333 32.5 36.6666 26.6667 36.6666H13.3333C7.5 36.6666 5 33.3333 5 28.3333V11.6666C5 6.66659 7.5 3.33325 13.3333 3.33325H26.6667C32.5 3.33325 35 6.66659 35 11.6666Z"
                      stroke="#0C1717"
                      strokeWidth="1.5"
                      strokeMiterlimit="10"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    />
                    <path
                      d="M24.167 7.5V10.8333C24.167 12.6667 25.667 14.1667 27.5003 14.1667H30.8337"
                      stroke="#0C1717"
                      strokeWidth="1.5"
                      strokeMiterlimit="10"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    />
                    <path
                      d="M13.333 21.6667H19.9997"
                      stroke="#0C1717"
                      strokeWidth="1.5"
                      strokeMiterlimit="10"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    />
                    <path
                      d="M13.333 28.3333H26.6663"
                      stroke="#0C1717"
                      strokeWidth="1.5"
                      strokeMiterlimit="10"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    />
                  </svg>
                  <div className="flex-1 text-left">
                    <p>{(fileOrFiles as File).name}</p>
                    <p className="text-gray-700">
                      {((fileOrFiles as File).size / (1024 * 1024)).toFixed(2)}
                      MB
                    </p>
                  </div>
                </div>
                <a
                  onClick={onClose}
                  className="absolute right-4 top-4 block h-4 w-4 cursor-pointer"
                  aria-disabled={fileOrFiles === null}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="17"
                    viewBox="0 0 16 17"
                    fill="none"
                  >
                    <path
                      d="M4.66602 12.079L11.3327 5.41235"
                      stroke="#0C1717"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    />
                    <path
                      d="M11.3327 12.079L4.66602 5.41235"
                      stroke="#0C1717"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    />
                  </svg>
                </a>
              </div>
            )}
          </div>
        )}
      </div>
    </FileUploader>
  );
}

export default Upload;
