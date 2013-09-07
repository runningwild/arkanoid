package main

type segment3 struct {
	p1 Point3
	p2 Point3
}

type plane3 struct {
	normal   Point3
	position float64
}

type collidablePolygon3 struct {
	polygon Polygon3
	plane   plane3
}

// Point vs plane face

func getPlaneCollision(p *plane3, s *segment3, r float64) *Point3 {
	v1 := s.p1.Dot(p.normal)
	v2 := s.p2.Dot(p.normal)
	f := 0.0
	if v1 > r && v2 <= r {
		f = (r - v1) / (v2 - v1)
	} else if v1 < -r && v2 >= -r {
		f = (-r - v1) / (v2 - v1)
	} else {
		return nil
	}
	result := s.p1.Times(f).Plus(s.p2.Times(1.0 - f))
	return &result
}

func isPointInsidePolygon(p *Point3, cp *collidablePolygon3) bool {
	hitPos := false
	hitNeg := false
	var points *[]Point3 = &cp.polygon.Points
	for i := 0; i < len(*points); i++ {
		j := (i + 1) % len(*points)
		side := (*points)[j].Minus((*points)[i])
		v := p.Minus((*points)[j])
		cross := v.Cross(side).Dot(cp.plane.normal)
		if cross > 0.0 {
			hitPos = true
		}
		if cross < 0.0 {
			hitNeg = true
		}
	}
	return !hitPos || !hitNeg
}

func getPolygonFaceCollision(
	s *segment3, r float64, cp *collidablePolygon3) *Point3 {
	p := getPlaneCollision(&cp.plane, s, r)
	if p != nil && isPointInsidePolygon(p, cp) {
		return p
	}
	return nil
}

// Data structure

func makeCollidablePolygon3(p *Polygon3) collidablePolygon3 {
	delta1 := p.Points[1].Minus(p.Points[0])
	delta2 := p.Points[2].Minus(p.Points[1])
	normal := delta1.Cross(delta2)
	normal = normal.Times(1 / normal.Length())
	plane := plane3{normal, normal.Dot(p.Points[0])}
	return collidablePolygon3{*p, plane}
}

type CollisionMesh3 struct {
	collidablePolygons []collidablePolygon3
}

func (cm *CollisionMesh3) AddPolygon(p *Polygon3) {
	cm.collidablePolygons = append(
		cm.collidablePolygons, makeCollidablePolygon3(p))
}
