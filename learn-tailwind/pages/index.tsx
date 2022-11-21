import Head from 'next/head'
import Image from 'next/image'
import Layout from '../components/Layout'

export default function Home() {
  return (
    <Layout home>
      <h3 className='text-4xl font-normal'>This is index</h3>
    </Layout>
  )
}
