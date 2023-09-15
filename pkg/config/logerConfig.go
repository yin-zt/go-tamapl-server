package config

var (
	LogCronConfigStr = `
<seelog type="asynctimer" asyncinterval="1000" minlevel="trace" maxlevel="error">
	<outputs formatid="common">
		<buffered formatid="common" size="1048576" flushperiod="1000">
			<rollingfile type="size" filename="./log/cron.log" maxsize="104857600" maxrolls="10"/>
		</buffered>
	</outputs>
	 <formats>
		 <format id="common" format="%Date %Time [%LEV] [%File:%Line] [%Func] %Msg%n" />
	 </formats>
</seelog>
`

	LogCmdConfigStr = `
<seelog type="asynctimer" asyncinterval="1000" minlevel="trace" maxlevel="error">
	<outputs formatid="common">
		<buffered formatid="common" size="1048576" flushperiod="1000">
			<rollingfile type="size" filename="./log/cmd.log" maxsize="104857600" maxrolls="10"/>
		</buffered>
	</outputs>
	 <formats>
		 <format id="common" format="%Date %Time [%LEV] [%File:%Line] [%Func] %Msg%n" />
	 </formats>
</seelog>
`

	LogServerConfigStr = `
<seelog type="asynctimer" asyncinterval="1000" minlevel="trace" maxlevel="error">
	<outputs formatid="common">
		<buffered formatid="common" size="1048576" flushperiod="1000">
			<rollingfile type="size" filename="./log/server.log" maxsize="104857600" maxrolls="10"/>
		</buffered>
	</outputs>
	 <formats>
		 <format id="common" format="%Date %Time [%LEV] [%File:%Line] [%Func] %Msg%n" />
	 </formats>
</seelog>
`

	LogAccessConfigStr = `
<seelog type="asynctimer" asyncinterval="1000" minlevel="trace" maxlevel="error">
	<outputs formatid="common">
		<buffered formatid="common" size="1048576" flushperiod="1000">
			<rollingfile type="size" filename="./log/access.log" maxsize="104857600" maxrolls="10"/>
		</buffered>
	</outputs>
	 <formats>
		 <format id="common" format="%Date %Time [%LEV] [%File:%Line] [%Func] %Msg%n" />
	 </formats>
</seelog>
`
)
