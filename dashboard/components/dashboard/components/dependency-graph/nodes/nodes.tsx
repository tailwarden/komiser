import React from 'react';
import { Handle, Position } from 'reactflow';
import Tippy from '@tippyjs/react';

type CircleNodeProps = {
  data: any;
  id: string;
};

function getInfo(data: any, id: any) {
  return (
    <>
      <div className="h-auto w-auto border bg-white">
        <div className="m-4">
          <p>Name: {data.label}</p>
          <p>ResourceId: {id}</p>
        </div>
      </div>
    </>
  );
}

function getLogo(label: string) {
  switch (label) {
    case 'EC2':
      return (
        <svg className="m-auto h-20 w-20" xmlns="http://www.w3.org/2000/svg">
          <defs>
            <linearGradient
              x1="0%"
              y1="100%"
              x2="100%"
              y2="0%"
              id="Arch_Amazon-EC2_32_svg__a"
            >
              <stop stopColor="#C8511B" offset="0%"></stop>
              <stop stopColor="#F90" offset="100%"></stop>
            </linearGradient>
          </defs>
          <g fill="none" fillRule="evenodd">
            <path
              d="M0 0h40v40H0z"
              fill="url(#Arch_Amazon-EC2_32_svg__a)"
            ></path>
            <path
              d="M26.052 27L26 13.948 13 14v13.052L26.052 27zM27 14h2v1h-2v2h2v1h-2v2h2v1h-2v2h2v1h-2v2h2v1h-2v.052a.95.95 0 01-.948.948H26v2h-1v-2h-2v2h-1v-2h-2v2h-1v-2h-2v2h-1v-2h-2v2h-1v-2h-.052a.95.95 0 01-.948-.948V27h-2v-1h2v-2h-2v-1h2v-2h-2v-1h2v-2h-2v-1h2v-2h-2v-1h2v-.052a.95.95 0 01.948-.948H13v-2h1v2h2v-2h1v2h2v-2h1v2h2v-2h1v2h2v-2h1v2h.052a.95.95 0 01.948.948V14zm-6 19H7V19h2v-1H7.062C6.477 18 6 18.477 6 19.062v13.876C6 33.523 6.477 34 7.062 34h13.877c.585 0 1.061-.477 1.061-1.062V31h-1v2zM34 7.062v13.876c0 .585-.476 1.062-1.061 1.062H30v-1h3V7H19v3h-1V7.062C18 6.477 18.477 6 19.062 6h13.877C33.524 6 34 6.477 34 7.062z"
              fill="#FFF"
            ></path>
          </g>
        </svg>
      );
    case 'Elastic IP':
      return (
        <svg className="m-auto h-20 w-20" xmlns="http://www.w3.org/2000/svg">
          <path
            d="M43.62 24.058v-.264l.122.132-.122.132zm-32.694 6.794A6.933 6.933 0 014 23.926 6.933 6.933 0 0110.926 17a6.933 6.933 0 016.925 6.926 6.932 6.932 0 01-6.925 6.926zm34.913-7.606l-6.441-6.936-1.465 1.361 4.88 5.255H19.792C19.293 18.474 15.51 15 10.926 15 6.004 15 2 19.004 2 23.926c0 4.922 4.004 8.926 8.926 8.926 4.584 0 8.367-3.474 8.866-7.926h23.021l-4.88 5.255 1.465 1.361 6.441-6.935a1 1 0 000-1.361z"
            fill="#D45B07"
            fillRule="evenodd"
          ></path>
        </svg>
      );
    case 'SUBNET':
      return <>SUBNET</>;
    case 'VPC':
      return <>VPC</>;
    case 'SECURITY GROUPS':
      return <>SECURITY GROUPS</>;
    case 'BLOCK DEVICE':
      return (
        <svg className="w-300 h-30 m-auto" xmlns="http://www.w3.org/2000/svg">
          <path
            d="M13.418 4H33.88l4.416 5.873H9.005L13.419 4zm-6.417 7.873h33.298a.998.998 0 00.895-.554.997.997 0 00-.096-1.047l-5.92-7.873A.998.998 0 0034.38 2H12.92a.998.998 0 00-.8.399l-5.918 7.873a.997.997 0 00-.096 1.047c.168.34.516.554.894.554zm1.06 32.112H39.3V15.371H8.06v28.614zM40.3 13.371H7.06a1 1 0 00-1 1v30.614a1 1 0 001 1H40.3a1 1 0 001-1V14.371a1 1 0 00-1-1z"
            fill="#3F8624"
            fillRule="evenodd"
          ></path>
        </svg>
      );
    case 'KEYPAIR':
      return <>KEYPAIR</>;
    default:
      return <></>
  }
}

const css = `
    .c {
        border: 1px solid black;
    }
`;
const CustomNode: React.FC<CircleNodeProps> = ({ data, id }) => (
  <>
    <style>{css}</style>
    <Tippy content={<>{getInfo(data, id)}</>}>
      <div className="c flex h-24 w-24 items-center justify-center rounded-[20%] border p-6">
        <Handle type="target" position={Position.Top} id={`${id}.top`} />
        <Handle type="source" position={Position.Bottom} id={`${id}.bottom`} />
        {getLogo(data.label)}
      </div>
    </Tippy>
  </>
);

export default CustomNode;
