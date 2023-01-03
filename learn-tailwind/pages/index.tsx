import Layout from '../components/Layout'
import Seo from '../components/SEO'

export default function Home() {
  return (
    <>
      <Seo/>
      <Layout home>
      {/* <div className='p-6 max-w-sm mx-auto bg-white rounded-xl shadow-lg flex items-center space-x-4'>
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
      </div> */}

      <div className='font-bebasneue'>
        <div id='index' className='flex items-center mb-96'>
          <div className='mx-80'>
            <h1 className='text-7xl'>FRANKIE TAYLOR SCOTT</h1>
            <h2 className='text-5xl'>Actress | Creator</h2>
          </div>
          <div className='flex items-center shadow-2xl'>
            <img src="/images/Frankie.png" alt="" />
          </div>
        </div>
        <div id='About' className='flex items-center justify-center h-screen space-x-4'>
          <div className='text-center'>
            <h2 className='text-4xl underline mb-10'>About Frankie</h2>
            <p className='text-left text-lg font-poppins max-w-4xl'>
              Frankie's pursuit of performance started from the age of 5 when she first started taking private singing classes and from this has branched into many creative avenues. 
              Taking her first singing examination with ABRSM at the age of 7, 
              Frankie continued her training and has since achieved Grade 8 Distinction in Musical Theatre performance. Frankie went on to attain her BA Hons Acting Degree studying at Bath School of Music and Performance graduating in 2021.
              Throughout her studies at university, Frankie discovered her passion for the art of voice over. 
              Since graduating she has had multiple classes training under notable talents and has continued to do so into her professional career. 
            </p>
            <img className='block mx-auto h-96 sm:shrink-0' src="/images/frankiestream.png" alt="" />
            <p className='max-w-4xl text-left text-lg font-poppins'>
            Frankie has also started her own successful streaming career on both YouTube and Twitch platforms, 
            achiving partnership on both. Frankie created her platforms to fuse together her love for video games and passion for performance. 
            Frankie, otherwise known as 'Frankie Lollia' on her platforms, 
            has grown her audience for the past 2 years amassing over 40,000 subscribers on YouTube for both her voice acting for 'The Otachan Show' as well as her condensed Video game playthroughs. 
            </p>
          </div>
        </div>
        <div id='Portfolio' className='flex items-center justify-center h-screen space-x-4'>
          <div className='text-center'>test</div>
        </div>
      </div>
      
    </Layout>
    </>
  )
}
