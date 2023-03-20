import { SVGProps } from 'react';

const DownloadIcon = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 24 24"
    fill="none"
    {...props}
  >
    <path
      stroke="currentColor"
      strokeLinecap="round"
      strokeLinejoin="round"
      strokeWidth="1.5"
      d="M9 11.5v6l2-2M9 17.5l-2-2"
    ></path>
    <path
      stroke="currentColor"
      strokeLinecap="round"
      strokeLinejoin="round"
      strokeWidth="1.5"
      d="M22 10.5v5c0 5-2 7-7 7H9c-5 0-7-2-7-7v-6c0-5 2-7 7-7h5"
    ></path>
    <path
      stroke="currentColor"
      strokeLinecap="round"
      strokeLinejoin="round"
      strokeWidth="1.5"
      d="M22 10.5h-4c-3 0-4-1-4-4v-4l8 8z"
    ></path>
  </svg>
);

export default DownloadIcon;
