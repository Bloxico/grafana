import { AccessControlAction, Team, TeamPermissionLevel } from 'app/types';
import { featureEnabled } from '@grafana/runtime';
import { NavModelItem, NavModel } from '@grafana/data';
import config from 'app/core/config';
import { ProBadge } from 'app/core/components/Upgrade/ProBadge';
import { contextSrv } from 'app/core/services/context_srv';

const loadingTeam = {
  avatarUrl: 'public/img/user_profile.png',
  id: 1,
  name: 'Loading',
  email: 'loading',
  memberCount: 0,
  permission: TeamPermissionLevel.Member,
};

export function buildNavModel(team: Team): NavModelItem {
  const navModel: NavModelItem = {
    img: team.avatarUrl,
    id: 'team-' + team.id,
    subTitle: 'Manage members and settings',
    url: '',
    text: team.name,
    breadcrumbs: [{ title: 'Teams', url: 'org/teams' }],
    children: [
      // With FGAC this tab will always be available (but not always editable)
      // With Legacy it will be hidden by hideTabsFromNonTeamAdmin should the user not be allowed to see it
      {
        active: false,
        icon: 'sliders-v-alt',
        id: `team-settings-${team.id}`,
        text: 'Settings',
        url: `org/teams/edit/${team.id}/settings`,
      },
    ],
  };

  // While team is loading we leave the members tab
  // With FGAC the Members tab is available when user has ActionTeamsPermissionsRead for this team
  // With Legacy it will always be present
  if (
    team === loadingTeam ||
    contextSrv.hasPermissionInMetadata(AccessControlAction.ActionTeamsPermissionsRead, team)
  ) {
    navModel.children!.unshift({
      active: false,
      icon: 'users-alt',
      id: `team-members-${team.id}`,
      text: 'Members',
      url: `org/teams/edit/${team.id}/members`,
    });
  }

  const teamGroupSync = {
    active: false,
    icon: 'sync',
    id: `team-groupsync-${team.id}`,
    text: 'External group sync',
    url: `org/teams/edit/${team.id}/groupsync`,
  };

  // With both Legacy and FGAC the tab is protected being featureEnabled
  // While team is loading we leave the teamsync tab
  // With FGAC the External Group Sync tab is available when user has ActionTeamsPermissionsRead for this team
  if (
    featureEnabled('teamsync') &&
    (team === loadingTeam || contextSrv.hasPermissionInMetadata(AccessControlAction.ActionTeamsPermissionsRead, team))
  ) {
    navModel.children!.push(teamGroupSync);
  } else if (config.featureToggles.featureHighlights) {
    navModel.children!.push({ ...teamGroupSync, tabSuffix: ProBadge });
  }

  return navModel;
}

export function getTeamLoadingNav(pageName: string): NavModel {
  const main = buildNavModel(loadingTeam);

  let node: NavModelItem;

  // find active page
  for (const child of main.children!) {
    if (child.id!.indexOf(pageName) > 0) {
      child.active = true;
      node = child;
      break;
    }
  }

  return {
    main: main,
    node: node!,
  };
}
