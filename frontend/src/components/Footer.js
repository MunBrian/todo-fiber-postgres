import React from 'react'
import { BsFillSuitHeartFill } from "react-icons/bs"

const Footer = () => {
  return (
      <footer className='flex items-center justify-center mt-3'>
      <h3 className='text-center text-white text-sm'>Made by <strong> Brian Mungai </strong> with love. </h3>
      <BsFillSuitHeartFill className='ml-2 fill-darkpurple' />
    </footer>
  )
}

export default Footer