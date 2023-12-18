import { ReactNode } from 'react';

export type DependencyGraphFilterFieldOptionsProps = {
  label: string;
  value: string;
  icon: ReactNode;
};

const DependencyGraphFilterFieldOptions: DependencyGraphFilterFieldOptionsProps[] =
  [
    {
      label: 'Cloud provider',
      value: 'provider',
      icon: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          fill="none"
          viewBox="0 0 24 24"
        >
          <path
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeMiterlimit="10"
            strokeWidth="1.5"
            d="M9 22H7c-4 0-5-1-5-5V7c0-4 1-5 5-5h1.5c1.5 0 1.83.44 2.4 1.2l1.5 2c.38.5.6.8 1.6.8h3c4 0 5 1 5 5v2"
          ></path>
          <path
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeMiterlimit="10"
            strokeWidth="1.5"
            d="M13.76 18.32c-2.35.17-2.35 3.57 0 3.74h5.56c.67 0 1.33-.25 1.82-.7 1.65-1.44.77-4.32-1.4-4.59-.78-4.69-7.56-2.91-5.96 1.56"
          ></path>
        </svg>
      )
    },
    {
      label: 'Cloud region',
      value: 'region',
      icon: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          fill="none"
          viewBox="0 0 24 24"
        >
          <path
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
            d="M22 12c0-5.52-4.48-10-10-10S2 6.48 2 12s4.48 10 10 10"
          ></path>
          <path
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
            d="M8 3h1a28.424 28.424 0 000 18H8M15 3c.97 2.92 1.46 5.96 1.46 9"
          ></path>
          <path
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
            d="M3 16v-1c2.92.97 5.96 1.46 9 1.46M3 9a28.424 28.424 0 0118 0M18.2 21.4a3.2 3.2 0 100-6.4 3.2 3.2 0 000 6.4zM22 22l-1-1"
          ></path>
        </svg>
      )
    },
    {
      label: 'Cloud service',
      value: 'service',
      icon: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          fill="none"
          viewBox="0 0 24 24"
        >
          <path
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeMiterlimit="10"
            strokeWidth="1.5"
            d="M6.37 9.51c-4.08.29-4.07 6.2 0 6.49h9.66c1.17.01 2.3-.43 3.17-1.22 2.86-2.5 1.33-7.5-2.44-7.98C15.41-1.34 3.62 1.75 6.41 9.51M12 16v3M12 23a2 2 0 100-4 2 2 0 000 4zM18 21h-4M10 21H6"
          ></path>
        </svg>
      )
    },
    {
      label: 'Resource relation',
      value: 'relations',
      icon: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          fill="none"
          viewBox="0 0 16 16"
        >
          <path
            d="M14.6667 10C14.6667 12.58 12.58 14.6667 10 14.6667L10.7 13.5"
            stroke="#697372"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M1.33334 5.9987C1.33334 3.4187 3.42 1.33203 6 1.33203L5.3 2.4987"
            stroke="#697372"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M9.13333 2.96484L11.7867 4.49817L14.4133 2.97152"
            stroke="#697372"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M11.7867 7.2122V4.49219"
            stroke="#697372"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M11.16 1.47203L9.56 2.35867C9.2 2.55867 8.9 3.06533 8.9 3.47866V5.17202C8.9 5.58536 9.19333 6.09202 9.56 6.29202L11.16 7.1787C11.5 7.37203 12.06 7.37203 12.4067 7.1787L14.0067 6.29202C14.3667 6.09202 14.6667 5.58536 14.6667 5.17202V3.47866C14.6667 3.06533 14.3733 2.55867 14.0067 2.35867L12.4067 1.47203C12.0667 1.28536 11.5067 1.28536 11.16 1.47203Z"
            stroke="#697372"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M1.56667 10.3008L4.21334 11.8341L6.84667 10.3075"
            stroke="#697372"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M4.21334 14.5481V11.8281"
            stroke="#697372"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M3.59334 8.80797L1.99334 9.69461C1.63334 9.89461 1.33334 10.4013 1.33334 10.8146V12.508C1.33334 12.9213 1.62667 13.428 1.99334 13.628L3.59334 14.5146C3.93334 14.708 4.49333 14.708 4.84 14.5146L6.44001 13.628C6.80001 13.428 7.1 12.9213 7.1 12.508V10.8146C7.1 10.4013 6.80667 9.89461 6.44001 9.69461L4.84 8.80797C4.49333 8.6213 3.93334 8.6213 3.59334 8.80797Z"
            stroke="#697372"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      )
    }
  ];

export default DependencyGraphFilterFieldOptions;
