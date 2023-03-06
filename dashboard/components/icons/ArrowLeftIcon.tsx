import { SVGProps } from 'react';

const ArrowLeftIcon = (props: SVGProps<SVGSVGElement>) => (
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
      strokeMiterlimit="10"
      strokeWidth="1.5"
      d="M9.57 18.07L3.5 12l6.07-6.07M20.5 12H3.67"
    ></path>
  </svg>
);

export default ArrowLeftIcon;
