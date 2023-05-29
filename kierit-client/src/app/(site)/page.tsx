import Image from 'next/image'
import AuthForm from './components/AuthForm'

export default function Home() {
  return (
    <div className='flex min-h-full flex-col justify-center py-12 sm:px-6 lg:px-8 bg-gray-100'>
      <div className='sm:mx-auto sm:w-full sm:max-w-md'>
        <div className='mx-auto w-auto flex justify-center'>
          <div className='flex '>
            <Image alt='logo' height='70' width='70' src='/images/bklogo.png'/>
            {/* <h3 className='ml-6 text-blue-800 font-bold text-2xl'> BACH KHOA Messenger</h3> */}
          </div>
        </div>
        <h2 className='mt-6 text-center text-2xl font-bold tracking-tight text-gray-900'>
          Sign in to your account
        </h2>
      </div>
      <AuthForm/>
    </div>
  )
}
