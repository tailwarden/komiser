import { CardProps } from './Card';

const base: CardProps = {
  label: 'Blop',
  value: 55,
  tooltip: 'You just hovered me! :O',
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
        d="M5.54 11.12c-4.68.33-4.68 7.14 0 7.47h1.92M5.59 11.12C2.38 2.19 15.92-1.38 17.47 8c4.33.55 6.08 6.32 2.8 9.19-1 .91-2.29 1.41-3.64 1.4h-.09"
      ></path>
      <path
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeMiterlimit="10"
        strokeWidth="1.5"
        d="M17 16.53c0 .74-.16 1.44-.46 2.06-.08.18-.17.35-.27.51A4.961 4.961 0 0112 21.53c-1.82 0-3.41-.98-4.27-2.43-.1-.16-.19-.33-.27-.51-.3-.62-.46-1.32-.46-2.06 0-2.76 2.24-5 5-5s5 2.24 5 5z"
      ></path>
      <path
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth="1.5"
        d="M10.44 16.53l.99.99 2.13-1.97"
      ></path>
    </svg>
  )
};

const mockCardProps = {
  base
};

export default mockCardProps;
