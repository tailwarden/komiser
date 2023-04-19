import { SVGProps } from 'react';

const StarIcon = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 16 16"
    fill="none"
    {...props}
  >
    <path
      stroke="currentColor"
      strokeLinecap="round"
      strokeLinejoin="round"
      strokeWidth="1.5"
      d="M9.153 2.34l1.174 2.346c.16.327.586.64.946.7l2.127.354c1.36.226 1.68 1.213.7 2.186L12.447 9.58c-.28.28-.434.82-.347 1.206l.473 2.047c.374 1.62-.486 2.247-1.92 1.4l-1.993-1.18c-.36-.213-.953-.213-1.32 0l-1.993 1.18c-1.427.847-2.294.213-1.92-1.4l.473-2.047c.087-.386-.067-.926-.347-1.206L1.9 7.926c-.973-.973-.66-1.96.7-2.186l2.127-.354c.353-.06.78-.373.94-.7L6.84 2.34c.64-1.274 1.68-1.274 2.313 0z"
    ></path>
  </svg>
);

export default StarIcon;
