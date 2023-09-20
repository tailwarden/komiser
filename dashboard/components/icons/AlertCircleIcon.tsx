import { SVGProps } from 'react';

const AlertCircleIcon = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 49 49"
    fill="none"
    {...props}
  >
    <circle cx="24.5" cy="24.5" r="24" fill="#FFE8E8" />
    <path
      d="M24.5 34.5C30.0228 34.5 34.5 30.0228 34.5 24.5C34.5 18.9772 30.0228 14.5 24.5 14.5C18.9772 14.5 14.5 18.9772 14.5 24.5C14.5 30.0228 18.9772 34.5 24.5 34.5Z"
      stroke="#DE5E5E"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
    <path
      d="M24.5 20.5V24.5"
      stroke="#DE5E5E"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
    <path
      d="M24.5 28.5H24.51"
      stroke="#DE5E5E"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
  </svg>
);

export default AlertCircleIcon;
