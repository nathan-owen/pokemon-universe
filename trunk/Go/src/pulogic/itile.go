/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package pulogic

import (
	list "container/list"
	pos "putools/pos"
)

const (
	TILEBLOCK_BLOCK       int = 1
	TILEBLOCK_WALK            = 2
	TILEBLOCK_SURF            = 3
	TILEBLOCK_TOP             = 4
	TILEBLOCK_BOTTOM          = 5
	TILEBLOCK_RIGHT           = 6
	TILEBLOCK_LEFT            = 7
	TILEBLOCK_TOPRIGHT        = 8
	TILEBLOCK_BOTTOMRIGHT     = 9
	TILEBLOCK_BOTTOMLEFT      = 10
	TILEBLOCK_TOPLEFT         = 11
)

type TilesMap map[int64]ITile
type LayerMap map[int]*TileLayer

type ITile interface {
	GetPosition()	pos.Position
	GetBlocking() 	int
	GetCreatures() 	CreatureList
	GetLayers()		LayerMap
	GetEvents() 	*list.List
	GetLocation()	ILocation
	AddCreature(_creature ICreature, _checkEvents bool) (ret int)
	RemoveCreature(_creature ICreature, _checkEvents bool) (ret int)
}

type TileLayer struct {
	Layer    int
	SpriteID int
}