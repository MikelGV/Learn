import Head from 'next/head'
import Image from 'next/image'
import Layout from '../components/Layout'

export default function Home() {
  return (
    <Layout home>
      <div className='p-6 max-w-sm mx-auto bg-white rounded-xl shadow-lg flex items-center space-x-4'>
        <div className='shrink-0'>
          <img className='h-12 w-12' src="/img/logo.svg" alt="ChitChat Logo" />
        </div>
        <div className='text-xl font-medium text-black'>ChitChat</div>
        <p className='text-slate-500'>You have a new message!</p>
      </div>

      <div className="py-8 px-8 max-w-sm mx-auto bg-white rounded-xl shadow-lg space-y-2 sm:py-4 sm:flex sm:items-center sm:space-y-0 sm:space-x-6">
        <img className="block mx-auto h-24 rounded-full sm:mx-0 sm:shrink-0" src="/images/Frankie.png" alt="Woman's Face"/>
        <div className="text-center space-y-2 sm:text-left">
          <div className="space-y-0.5">
            <p className="text-lg text-black font-semibold">
              Erin Lindford
            </p>
            <p className="text-slate-500 font-medium">
              Product Engineer
            </p>
          </div>
          <button className="px-4 py-1 text-sm text-purple-600 font-semibold rounded-full border border-purple-200 hover:text-white hover:bg-purple-600 hover:border-transparent focus:outline-none focus:ring-2 focus:ring-purple-600 focus:ring-offset-2">Message</button>
        </div>
      </div>

      <div className='p-6 font-bebasneue flex items-center'>
        <div className='mx-80'>
          <h1 className='text-7xl'>FRANKIE TAYLOR SCOTT</h1>
          <h2 className='text-5xl'>Actress | Creator</h2>
        </div>
        <div>
          <img className='' src="/images/Frankie.png" alt="" />
        </div>
      </div>
      
    </Layout>
  )
}
