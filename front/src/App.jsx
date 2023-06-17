import { Fragment, useEffect, useState } from "react"
import Cookies from 'universal-cookie'
import useSound from 'use-sound'
import Click from './assets/click.mp3'
import Button from "./components/button"
import DesktopButton from "./components/desktop/button"
import DockPanel from "./components/dockpanel"
import Modal from "./components/modal"
import TaskbarButton from "./components/taskbar/button"
import Clock from "./components/taskbar/clock"
import { useMediaQuery } from "./utils/useMediaQuery"

const App = () => {

  const isMobile = useMediaQuery(480)

  const [play] = useSound(Click, {
    volume: 0.1
  })
  
  const [dockpanel, setDockpanel] = useState(false)
  const [connectWallet, setConnectWallet] = useState(false)
  const [selectWallet, setSelectWallet] = useState('')
  const [about, setAbout] = useState(false)
  const [auth, setAuth] = useState(null)

  const cookies = new Cookies()
  const jwt = cookies.get('jwt');

  useEffect(() => {
    
    const queryString = window.location.search
    const searchParams = new URLSearchParams(queryString)

    if (searchParams.has('jwt')) {
      const jwt = searchParams.get('jwt')

      cookies.set('jwt', jwt, { path: '/' });

      searchParams.delete('jwt')
      window.location.search = searchParams

    } 

    checkAuth()

  }, [])

  const handleAuth = () => window.open('https://back.punkzhub.com/api/auth?provider=discord')

  const checkAuth = async () => {
    if (jwt) {
      fetch('https://back.punkzhub.com/api/user', {
        method: 'GET', 
        headers: {
          "Authorization" : jwt
        }
      })
      .then(async (res) => {
        if (res.ok) {
          setAuth(await res.json())
        }
      })
    }
  }

  const sendWallet = () => {
    if (selectWallet && jwt) {
      fetch('https://back.punkzhub.com/api/wallet', {
        method: 'POST',
        headers: {
          "Authorization" : jwt
        },
        body: JSON.stringify({wallet: selectWallet})
      }).then(async (res) => {
          if (res.ok) {
            setSelectWallet('')
            alert('Awesome! Your wallet is registered!')
          } else {
            const { error } = await res.json()
            error ? alert(error) : alert('Bad request!')
          }
      })
    }
  }

  return (
    <div className="relative overflow-hidden h-full bg-body bg-cover bg-center flex flex-col font-segoe z-10" onClick={play}>
        {!(isMobile) && (
          <Fragment>
            {/* Desktop */}
            <div className="h-full relative" onClick={() => setDockpanel(false)}>
              <DesktopButton hoverable hover={connectWallet} icon={'bg-computer'} title={'PUNKZ!!1'} className={'top-[40px] left-[40px]'} onClick={() => setConnectWallet(prevState => !prevState)} />
              <DesktopButton hoverable icon={'bg-recycle'} title={'Recycle bin'} className={'bottom-[40px] right-[40px]'} onClick={() => window.open('https://punkz.gitbook.io/punkz/', '_blank')} />
            </div>

            {/* Task bar */}
            <div className="flex w-full radial-gradient z-30">
              <div className="flex w-full h-[41px] backdrop-blur-[3px] border-t-2 border-l-2 border-b-2 border-solid border-[#ffffff33]">
                {/* Win button */}
                <div className="bg-win-inactive hover:bg-win-active bg-center bg-no-repeat bg-contain min-h-[39px] h-[39px] min-w-[39px] w-[39px] cursor-pointer ml-[6px] mr-[12px]" onClick={() => setDockpanel(prevState => !prevState)} />
                <div className="flex gap-[2px] w-full">
                  {/* Twitter button */}
                  <TaskbarButton icon={'bg-twitter-logo'} onClick={() => window.open('https://twitter.com/wepunkz', '_blank')} />
                  {/* Discord button */}
                  <TaskbarButton icon={'bg-discord-logo'} onClick={() => window.open('https://discord.gg/wepunkz', '_blank')} />
                </div>
                {/* Clock */}
                <div className="flex relative h-full mr-[8px]">
                  <Clock/>
                </div>
              </div>
              {/* Hide all */}
              <div className="w-[15px] h-full border-[1px] border-solid border-[#00000080] shadow-hideButton button-hideall-gradient cursor-pointer" onClick={() => {
                setAbout(false)
                setDockpanel(false)
                setConnectWallet(false)
              }} />
            </div>
          </Fragment>
        )}

        <div className={isMobile ? 'flex flex-col absolute left-0 top-0 w-full h-full justify-center gap-[40px]' : ''}>
          {/* Connect Wallet Modal */}
          <Modal icon={'bg-computer'} title={'PUNKZ!!1'} isOpen={connectWallet} onClose={() => setConnectWallet(false)} className={`${(isMobile && about) ? 'relative !top-0' : ''}`}>
            <Modal.Body>
              <div className="flex flex-col items-center w-full">
                  <div className="text-[#40529C] text-[25px] leading-[33px]">Connect Wallet</div>
                  {(auth) && (
                    <div className="whitespace-nowrap overflow-hidden text-ellipsis max-w-[300px]">Your Discord: {auth.name}</div>
                  )}
                  {auth ? (
                    <Fragment>
                      <div className="max-w-[316px] w-full modal-input-border mt-[14px]">
                        <input className={`w-full bg-[#ECF2F9] rounded-[2px] modal-input px-[12px] ${isMobile ? 'h-[40px]' : 'h-[30px]'}`} type="text" placeholder="Wallet Address" value={selectWallet} onChange={(e) => setSelectWallet(e.target.value)} />
                      </div>
                      <div className="mt-[10px]">
                          <Button onClick={sendWallet}>Accept</Button>
                      </div>
                    </Fragment>
                  ) : (
                    <Fragment>
                      <div className="bg-[#6c88e0] rounded-[6px] text-white px-[20px] py-[10px] cursor-pointer mt-[20px] flex items-center" onClick={handleAuth}>
                        <div className="bg-discord-button bg-contain bg-no-repeat bg-center w-[40px] h-[40px]" style={{filter: 'drop-shadow(0px 0px 1px #00000085)'}} />
                        <b className="ml-[8px]">Log In with Discord</b>
                      </div>
                    </Fragment>
                  )}
              </div>
            </Modal.Body>
          </Modal>
          {/* About PUNKZ Modal */}
          <Modal icon={'bg-help'} title={'About PUNKZ'} isOpen={about && isMobile} onClose={() => setAbout(false)}>
            <Modal.Body>
              <div className="px-[30px] text-[15px] tracking-[-0.05em]">
              PUNKZ is an ecosystem of NFT utilities. The main utility is Omnistep, it is a platform to comfortably manage your project and to keep all your data and tasks in a simple way. The main utility of our platform is that we become a guarantor of work done and money paid.
              </div>
            </Modal.Body>
          </Modal>
        </div>
        <DockPanel open={isMobile || dockpanel} isMobile={isMobile} actions={{connectWallet, setConnectWallet, about, setAbout, auth}} />
    </div>
  )
}

export default App