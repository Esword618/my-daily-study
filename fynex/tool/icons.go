package tool

import (
	_ "embed"
	_ "unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	octicons "github.com/lusingander/fyne-octicons"
)

// https://blog.csdn.net/gold0523/article/details/94341008
//go:linkname icons fyne.io/fyne/v2/theme.icons
var icons map[fyne.ThemeIconName]fyne.Resource

func iconScreen() {
	txt := widget.NewEntry()
	c := container.NewGridWrap(fyne.NewSize(50, 50))
	cc := container.NewVScroll(container.NewVBox(txt, c))
	for _, icon := range icons {
		//c.Add(widget.NewIcon(icon))
		btn := widget.NewButtonWithIcon("", icon, nil)
		btn.OnTapped = func() {
			// fmt.Println(btn.Icon.Name())
			txt.SetText(btn.Icon.Name())
		}

		c.Add(btn)
	}
	// for _, icon := range icons2 {
	// 	btn := widget.NewButtonWithIcon("", icon.resource, nil)
	// 	btn.OnTapped = func() {
	// 		txt.SetText(btn.Icon.Name())
	// 	}

	// 	c.Add(btn)
	// }
	w := fyne.CurrentApp().NewWindow("Icons")
	w.Resize(fyne.NewSize(800, 700))
	w.SetContent(cc)
	w.Show()
}

var (
	icons1 = []fyne.Resource{
		octicons.AlertIcon(),
		octicons.BellIcon(),
		octicons.BoldIcon(),
		octicons.BookIcon(),
		octicons.BookmarkIcon(),
		octicons.BriefcaseIcon(),
		octicons.BroadcastIcon(),
		octicons.BrowserIcon(),
		octicons.BuIcon(),
		octicons.CalendarIcon(),
		octicons.CheckIcon(),
		octicons.ChecklistIcon(),
		octicons.ChevrondownIcon(),
		octicons.ChevronleftIcon(),
		octicons.ChevronrightIcon(),
		octicons.ChevronupIcon(),
		octicons.CircleslashIcon(),
		octicons.CircuitboardIcon(),
		octicons.ClippyIcon(),
		octicons.ClockIcon(),
		octicons.ClouddownloadIcon(),
		octicons.ClouduploadIcon(),
		octicons.CodeIcon(),
		octicons.CommentdiscussionIcon(),
		octicons.CommentIcon(),
		octicons.CreditcardIcon(),
		octicons.DashIcon(),
		octicons.DashboardIcon(),
		octicons.DatabaseIcon(),
		octicons.DependentIcon(),
		octicons.DesktopdownloadIcon(),
		octicons.DevicecameravideoIcon(),
		octicons.DevicecameraIcon(),
		octicons.DevicedesktopIcon(),
		octicons.DevicemobileIcon(),
		octicons.DiffaddedIcon(),
		octicons.DiffignoredIcon(),
		octicons.DiffmodifiedIcon(),
		octicons.DiffremovedIcon(),
		octicons.DiffrenamedIcon(),
		octicons.DiffIcon(),
		octicons.EllipsisIcon(),
		octicons.EyeclosedIcon(),
		octicons.EyeIcon(),
		octicons.FilebinaryIcon(),
		octicons.FilecodeIcon(),
		octicons.FiledirectoryIcon(),
		octicons.FilemediaIcon(),
		octicons.FilepdfIcon(),
		octicons.FilesubmoduleIcon(),
		octicons.FilesymlinkdirectoryIcon(),
		octicons.FilesymlinkfileIcon(),
		octicons.FilezipIcon(),
		octicons.FileIcon(),
		octicons.FlameIcon(),
		octicons.FolddownIcon(),
		octicons.FoldupIcon(),
		octicons.FoldIcon(),
		octicons.GearIcon(),
		octicons.GiftIcon(),
		octicons.GistsecretIcon(),
		octicons.GistIcon(),
		octicons.GitbranchIcon(),
		octicons.GitcommitIcon(),
		octicons.GitcompareIcon(),
		octicons.GitmergeIcon(),
		octicons.GitpullrequestIcon(),
		octicons.GithubactionIcon(),
		octicons.GlobeIcon(),
		octicons.GrabberIcon(),
		octicons.GraphIcon(),
		octicons.HeartoutlineIcon(),
		octicons.HeartIcon(),
		octicons.HistoryIcon(),
		octicons.HomeIcon(),
		octicons.HorizontalruleIcon(),
		octicons.HubotIcon(),
		octicons.InboxIcon(),
		octicons.InfinityIcon(),
		octicons.InfoIcon(),
		octicons.InternalrepoIcon(),
		octicons.IssueclosedIcon(),
		octicons.IssueopenedIcon(),
		octicons.IssuereopenedIcon(),
		octicons.ItalicIcon(),
		octicons.JerseyIcon(),
		octicons.KebabhorizontalIcon(),
		octicons.KebabverticalIcon(),
		octicons.KeyIcon(),
		octicons.KeyboardIcon(),
		octicons.LawIcon(),
		octicons.LightbulbIcon(),
		octicons.LinearrowdownIcon(),
		octicons.LinearrowleftIcon(),
		octicons.LinearrowrightIcon(),
		octicons.LinearrowupIcon(),
		octicons.LinkexternalIcon(),
		octicons.LinkIcon(),
		octicons.ListorderedIcon(),
		octicons.ListunorderedIcon(),
		octicons.LocationIcon(),
		octicons.LockIcon(),
		octicons.LogogistIcon(),
		octicons.LogogithubIcon(),
		octicons.MailreadIcon(),
		octicons.MailIcon(),
		octicons.MarkgithubIcon(),
		octicons.MarkdownIcon(),
		octicons.MegaphoneIcon(),
		octicons.MentionIcon(),
		octicons.MilestoneIcon(),
		octicons.MirrorIcon(),
		octicons.MortarboardIcon(),
		octicons.MuteIcon(),
		octicons.NonewlineIcon(),
		octicons.NorthstarIcon(),
		octicons.NoteIcon(),
		octicons.OctofaceIcon(),
		octicons.OrganizationIcon(),
		octicons.PackageIcon(),
		octicons.PaintcanIcon(),
		octicons.PencilIcon(),
		octicons.PersonIcon(),
		octicons.PinIcon(),
		octicons.PlayIcon(),
		octicons.PluIcon(),
		octicons.PlussmallIcon(),
		octicons.PlusIcon(),
		octicons.PrimitivedotstrokeIcon(),
		octicons.PrimitivedotIcon(),
		octicons.PrimitivesquareIcon(),
		octicons.ProjectIcon(),
		octicons.PulseIcon(),
		octicons.QuestionIcon(),
		octicons.QuoteIcon(),
		octicons.RadiotowerIcon(),
		octicons.ReplyIcon(),
		octicons.RepocloneIcon(),
		octicons.RepoforcepushIcon(),
		octicons.RepoforkedIcon(),
		octicons.RepopullIcon(),
		octicons.RepopushIcon(),
		octicons.RepotemplateprivateIcon(),
		octicons.RepotemplateIcon(),
		octicons.RepoIcon(),
		octicons.ReportIcon(),
		octicons.RequestchangesIcon(),
		octicons.RocketIcon(),
		octicons.RssIcon(),
		octicons.RubyIcon(),
		octicons.SavedIcon(),
		octicons.ScreenfullIcon(),
		octicons.ScreennormalIcon(),
		octicons.SearchIcon(),
		octicons.ServerIcon(),
		octicons.SettingsIcon(),
		octicons.ShieldcheckIcon(),
		octicons.ShieldlockIcon(),
		octicons.ShieldxIcon(),
		octicons.ShieldIcon(),
		octicons.SigninIcon(),
		octicons.SignoutIcon(),
		octicons.SkipIcon(),
		octicons.SmileyIcon(),
		octicons.SquirrelIcon(),
		octicons.StarIcon(),
		octicons.StopIcon(),
		octicons.SyncIcon(),
		octicons.TaIcon(),
		octicons.TasklistIcon(),
		octicons.TelescopeIcon(),
		octicons.TerminalIcon(),
		octicons.TextsizeIcon(),
		octicons.ThreebarsIcon(),
		octicons.ThumbsdownIcon(),
		octicons.ThumbsupIcon(),
		octicons.ToolsIcon(),
		octicons.TrashcanIcon(),
		octicons.TriangledownIcon(),
		octicons.TriangleleftIcon(),
		octicons.TrianglerightIcon(),
		octicons.TriangleupIcon(),
		octicons.UnfoldIcon(),
		octicons.UnmuteIcon(),
		octicons.UnsavedIcon(),
		octicons.UnverifiedIcon(),
		octicons.VerifiedIcon(),
		octicons.VersionsIcon(),
		octicons.WatchIcon(),
		octicons.WorkflowallIcon(),
		octicons.WorkflowIcon(),
		octicons.XIcon(),
		octicons.ZapIcon(),
	}
)

// https://github.com/lusingander/fyne-octicons/blob/master/internal/cmd/fyne-octicons/icons.go
var icons2 = []struct {
	name     string
	resource fyne.Resource
}{
	{"AlertIcon", octicons.AlertIcon()},
	{"ArchiveIcon", octicons.ArchiveIcon()},
	{"ArrowbothIcon", octicons.ArrowbothIcon()},
	{"ArrowdownIcon", octicons.ArrowdownIcon()},
	{"ArrowleftIcon", octicons.ArrowleftIcon()},
	{"ArrowrightIcon", octicons.ArrowrightIcon()},
	{"ArrowsmalldownIcon", octicons.ArrowsmalldownIcon()},
	{"ArrowsmallleftIcon", octicons.ArrowsmallleftIcon()},
	{"ArrowsmallrightIcon", octicons.ArrowsmallrightIcon()},
	{"ArrowsmallupIcon", octicons.ArrowsmallupIcon()},
	{"ArrowupIcon", octicons.ArrowupIcon()},
	{"BeakerIcon", octicons.BeakerIcon()},
	{"BellIcon", octicons.BellIcon()},
	{"BoldIcon", octicons.BoldIcon()},
	{"BookIcon", octicons.BookIcon()},
	{"BookmarkIcon", octicons.BookmarkIcon()},
	{"BriefcaseIcon", octicons.BriefcaseIcon()},
	{"BroadcastIcon", octicons.BroadcastIcon()},
	{"BrowserIcon", octicons.BrowserIcon()},
	{"BuIcon", octicons.BuIcon()},
	{"CalendarIcon", octicons.CalendarIcon()},
	{"CheckIcon", octicons.CheckIcon()},
	{"ChecklistIcon", octicons.ChecklistIcon()},
	{"ChevrondownIcon", octicons.ChevrondownIcon()},
	{"ChevronleftIcon", octicons.ChevronleftIcon()},
	{"ChevronrightIcon", octicons.ChevronrightIcon()},
	{"ChevronupIcon", octicons.ChevronupIcon()},
	{"CircleslashIcon", octicons.CircleslashIcon()},
	{"CircuitboardIcon", octicons.CircuitboardIcon()},
	{"ClippyIcon", octicons.ClippyIcon()},
	{"ClockIcon", octicons.ClockIcon()},
	{"ClouddownloadIcon", octicons.ClouddownloadIcon()},
	{"ClouduploadIcon", octicons.ClouduploadIcon()},
	{"CodeIcon", octicons.CodeIcon()},
	{"CommentdiscussionIcon", octicons.CommentdiscussionIcon()},
	{"CommentIcon", octicons.CommentIcon()},
	{"CreditcardIcon", octicons.CreditcardIcon()},
	{"DashIcon", octicons.DashIcon()},
	{"DashboardIcon", octicons.DashboardIcon()},
	{"DatabaseIcon", octicons.DatabaseIcon()},
	{"DependentIcon", octicons.DependentIcon()},
	{"DesktopdownloadIcon", octicons.DesktopdownloadIcon()},
	{"DevicecameravideoIcon", octicons.DevicecameravideoIcon()},
	{"DevicecameraIcon", octicons.DevicecameraIcon()},
	{"DevicedesktopIcon", octicons.DevicedesktopIcon()},
	{"DevicemobileIcon", octicons.DevicemobileIcon()},
	{"DiffaddedIcon", octicons.DiffaddedIcon()},
	{"DiffignoredIcon", octicons.DiffignoredIcon()},
	{"DiffmodifiedIcon", octicons.DiffmodifiedIcon()},
	{"DiffremovedIcon", octicons.DiffremovedIcon()},
	{"DiffrenamedIcon", octicons.DiffrenamedIcon()},
	{"DiffIcon", octicons.DiffIcon()},
	{"EllipsisIcon", octicons.EllipsisIcon()},
	{"EyeclosedIcon", octicons.EyeclosedIcon()},
	{"EyeIcon", octicons.EyeIcon()},
	{"FilebinaryIcon", octicons.FilebinaryIcon()},
	{"FilecodeIcon", octicons.FilecodeIcon()},
	{"FiledirectoryIcon", octicons.FiledirectoryIcon()},
	{"FilemediaIcon", octicons.FilemediaIcon()},
	{"FilepdfIcon", octicons.FilepdfIcon()},
	{"FilesubmoduleIcon", octicons.FilesubmoduleIcon()},
	{"FilesymlinkdirectoryIcon", octicons.FilesymlinkdirectoryIcon()},
	{"FilesymlinkfileIcon", octicons.FilesymlinkfileIcon()},
	{"FilezipIcon", octicons.FilezipIcon()},
	{"FileIcon", octicons.FileIcon()},
	{"FlameIcon", octicons.FlameIcon()},
	{"FolddownIcon", octicons.FolddownIcon()},
	{"FoldupIcon", octicons.FoldupIcon()},
	{"FoldIcon", octicons.FoldIcon()},
	{"GearIcon", octicons.GearIcon()},
	{"GiftIcon", octicons.GiftIcon()},
	{"GistsecretIcon", octicons.GistsecretIcon()},
	{"GistIcon", octicons.GistIcon()},
	{"GitbranchIcon", octicons.GitbranchIcon()},
	{"GitcommitIcon", octicons.GitcommitIcon()},
	{"GitcompareIcon", octicons.GitcompareIcon()},
	{"GitmergeIcon", octicons.GitmergeIcon()},
	{"GitpullrequestIcon", octicons.GitpullrequestIcon()},
	{"GithubactionIcon", octicons.GithubactionIcon()},
	{"GlobeIcon", octicons.GlobeIcon()},
	{"GrabberIcon", octicons.GrabberIcon()},
	{"GraphIcon", octicons.GraphIcon()},
	{"HeartoutlineIcon", octicons.HeartoutlineIcon()},
	{"HeartIcon", octicons.HeartIcon()},
	{"HistoryIcon", octicons.HistoryIcon()},
	{"HomeIcon", octicons.HomeIcon()},
	{"HorizontalruleIcon", octicons.HorizontalruleIcon()},
	{"HubotIcon", octicons.HubotIcon()},
	{"InboxIcon", octicons.InboxIcon()},
	{"InfinityIcon", octicons.InfinityIcon()},
	{"InfoIcon", octicons.InfoIcon()},
	{"InternalrepoIcon", octicons.InternalrepoIcon()},
	{"IssueclosedIcon", octicons.IssueclosedIcon()},
	{"IssueopenedIcon", octicons.IssueopenedIcon()},
	{"IssuereopenedIcon", octicons.IssuereopenedIcon()},
	{"ItalicIcon", octicons.ItalicIcon()},
	{"JerseyIcon", octicons.JerseyIcon()},
	{"KebabhorizontalIcon", octicons.KebabhorizontalIcon()},
	{"KebabverticalIcon", octicons.KebabverticalIcon()},
	{"KeyIcon", octicons.KeyIcon()},
	{"KeyboardIcon", octicons.KeyboardIcon()},
	{"LawIcon", octicons.LawIcon()},
	{"LightbulbIcon", octicons.LightbulbIcon()},
	{"LinearrowdownIcon", octicons.LinearrowdownIcon()},
	{"LinearrowleftIcon", octicons.LinearrowleftIcon()},
	{"LinearrowrightIcon", octicons.LinearrowrightIcon()},
	{"LinearrowupIcon", octicons.LinearrowupIcon()},
	{"LinkexternalIcon", octicons.LinkexternalIcon()},
	{"LinkIcon", octicons.LinkIcon()},
	{"ListorderedIcon", octicons.ListorderedIcon()},
	{"ListunorderedIcon", octicons.ListunorderedIcon()},
	{"LocationIcon", octicons.LocationIcon()},
	{"LockIcon", octicons.LockIcon()},
	{"LogogistIcon", octicons.LogogistIcon()},
	{"LogogithubIcon", octicons.LogogithubIcon()},
	{"MailreadIcon", octicons.MailreadIcon()},
	{"MailIcon", octicons.MailIcon()},
	{"MarkgithubIcon", octicons.MarkgithubIcon()},
	{"MarkdownIcon", octicons.MarkdownIcon()},
	{"MegaphoneIcon", octicons.MegaphoneIcon()},
	{"MentionIcon", octicons.MentionIcon()},
	{"MilestoneIcon", octicons.MilestoneIcon()},
	{"MirrorIcon", octicons.MirrorIcon()},
	{"MortarboardIcon", octicons.MortarboardIcon()},
	{"MuteIcon", octicons.MuteIcon()},
	{"NonewlineIcon", octicons.NonewlineIcon()},
	{"NorthstarIcon", octicons.NorthstarIcon()},
	{"NoteIcon", octicons.NoteIcon()},
	{"OctofaceIcon", octicons.OctofaceIcon()},
	{"OrganizationIcon", octicons.OrganizationIcon()},
	{"PackageIcon", octicons.PackageIcon()},
	{"PaintcanIcon", octicons.PaintcanIcon()},
	{"PencilIcon", octicons.PencilIcon()},
	{"PersonIcon", octicons.PersonIcon()},
	{"PinIcon", octicons.PinIcon()},
	{"PlayIcon", octicons.PlayIcon()},
	{"PluIcon", octicons.PluIcon()},
	{"PlussmallIcon", octicons.PlussmallIcon()},
	{"PlusIcon", octicons.PlusIcon()},
	{"PrimitivedotstrokeIcon", octicons.PrimitivedotstrokeIcon()},
	{"PrimitivedotIcon", octicons.PrimitivedotIcon()},
	{"PrimitivesquareIcon", octicons.PrimitivesquareIcon()},
	{"ProjectIcon", octicons.ProjectIcon()},
	{"PulseIcon", octicons.PulseIcon()},
	{"QuestionIcon", octicons.QuestionIcon()},
	{"QuoteIcon", octicons.QuoteIcon()},
	{"RadiotowerIcon", octicons.RadiotowerIcon()},
	{"ReplyIcon", octicons.ReplyIcon()},
	{"RepocloneIcon", octicons.RepocloneIcon()},
	{"RepoforcepushIcon", octicons.RepoforcepushIcon()},
	{"RepoforkedIcon", octicons.RepoforkedIcon()},
	{"RepopullIcon", octicons.RepopullIcon()},
	{"RepopushIcon", octicons.RepopushIcon()},
	{"RepotemplateprivateIcon", octicons.RepotemplateprivateIcon()},
	{"RepotemplateIcon", octicons.RepotemplateIcon()},
	{"RepoIcon", octicons.RepoIcon()},
	{"ReportIcon", octicons.ReportIcon()},
	{"RequestchangesIcon", octicons.RequestchangesIcon()},
	{"RocketIcon", octicons.RocketIcon()},
	{"RssIcon", octicons.RssIcon()},
	{"RubyIcon", octicons.RubyIcon()},
	{"SavedIcon", octicons.SavedIcon()},
	{"ScreenfullIcon", octicons.ScreenfullIcon()},
	{"ScreennormalIcon", octicons.ScreennormalIcon()},
	{"SearchIcon", octicons.SearchIcon()},
	{"ServerIcon", octicons.ServerIcon()},
	{"SettingsIcon", octicons.SettingsIcon()},
	{"ShieldcheckIcon", octicons.ShieldcheckIcon()},
	{"ShieldlockIcon", octicons.ShieldlockIcon()},
	{"ShieldxIcon", octicons.ShieldxIcon()},
	{"ShieldIcon", octicons.ShieldIcon()},
	{"SigninIcon", octicons.SigninIcon()},
	{"SignoutIcon", octicons.SignoutIcon()},
	{"SkipIcon", octicons.SkipIcon()},
	{"SmileyIcon", octicons.SmileyIcon()},
	{"SquirrelIcon", octicons.SquirrelIcon()},
	{"StarIcon", octicons.StarIcon()},
	{"StopIcon", octicons.StopIcon()},
	{"SyncIcon", octicons.SyncIcon()},
	{"TaIcon", octicons.TaIcon()},
	{"TasklistIcon", octicons.TasklistIcon()},
	{"TelescopeIcon", octicons.TelescopeIcon()},
	{"TerminalIcon", octicons.TerminalIcon()},
	{"TextsizeIcon", octicons.TextsizeIcon()},
	{"ThreebarsIcon", octicons.ThreebarsIcon()},
	{"ThumbsdownIcon", octicons.ThumbsdownIcon()},
	{"ThumbsupIcon", octicons.ThumbsupIcon()},
	{"ToolsIcon", octicons.ToolsIcon()},
	{"TrashcanIcon", octicons.TrashcanIcon()},
	{"TriangledownIcon", octicons.TriangledownIcon()},
	{"TriangleleftIcon", octicons.TriangleleftIcon()},
	{"TrianglerightIcon", octicons.TrianglerightIcon()},
	{"TriangleupIcon", octicons.TriangleupIcon()},
	{"UnfoldIcon", octicons.UnfoldIcon()},
	{"UnmuteIcon", octicons.UnmuteIcon()},
	{"UnsavedIcon", octicons.UnsavedIcon()},
	{"UnverifiedIcon", octicons.UnverifiedIcon()},
	{"VerifiedIcon", octicons.VerifiedIcon()},
	{"VersionsIcon", octicons.VersionsIcon()},
	{"WatchIcon", octicons.WatchIcon()},
	{"WorkflowallIcon", octicons.WorkflowallIcon()},
	{"WorkflowIcon", octicons.WorkflowIcon()},
	{"XIcon", octicons.XIcon()},
	{"ZapIcon", octicons.ZapIcon()},
}
